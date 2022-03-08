import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';
import { Property, Schema } from 'src/app/shared/models/schema.model';
import { ConstraintControlChange } from '../../components/constraint-control/constraint-control.component';

@Component({
  selector: 'app-constraint',
  templateUrl: './constraint.component.html',
  styleUrls: ['./constraint.component.scss']
})
export class ConstraintComponent implements OnInit {
  fetching = true;
  schemas: Schema[];
  schemasFiltered: Schema[];
  singleSchema: Schema;
  singleSchemaFiltered: Schema;
  singleSchemaTitle: string;

  constructor(private validationApiService: ValidationApiService, private route: ActivatedRoute) {}

  ngOnInit(): void {
    const segments = this.route.snapshot.url.map(segment => segment.path);
    this.singleSchemaTitle = segments[segments.length - 1];
    const constraintPath = segments.join('/');

    this.validationApiService.getConstraintForPath(constraintPath).subscribe(schemas => {
      this.fetching = false;
      if (schemas instanceof Array) {
        this.schemas = schemas;
        this.schemasFiltered = this.schemas.sort(this.groupKindVersionSorter());
      } else {
        this.singleSchema = schemas;
        this.singleSchemaFiltered = this.singleSchema;
      }
    });
  }

  getGroupKindVersionName(schema: Schema): string {
    const { group, kind, version } = schema['x-kubernetes-group-version-kind'][0];

    let groupString = '';
    if (group !== '') {
      groupString = '-' + group;
    }

    return kind + groupString + '-' + version;
  }

  getSingleSchemaPropertyLength(): number {
    return Object.keys(this.singleSchema.properties).length;
  }

  onSchemaDisableToggled(schemas: Schema[], index: number, disabledValue: boolean): void {
    const schema = { ...schemas[index] };

    if (schema['x-constraint'] === null) {
      schema['x-constraint'] = {};
    }
    schema['x-constraint'].disabled = disabledValue;

    schemas[index] = schema;
  }

  onSingleSchemaDisableToggled(
    schemaProperties: Record<string, Property>,
    keyName: string,
    disabledValue: boolean
  ): void {
    const property = { ...schemaProperties[keyName] };

    if (property['x-constraint'] === null) {
      property['x-constraint'] = {};
    }
    property['x-constraint'].disabled = disabledValue;

    schemaProperties[keyName] = property;
  }

  onSchemaListControlChange(changes: ConstraintControlChange): void {
    this.schemasFiltered = this.schemas
      // Filter deleted items
      .filter(schema => {
        if (changes.hideDeleted) {
          return !schema['x-constraint']?.disabled;
        }
        return true;
      })
      // Filter search
      .filter(schema =>
        this.getGroupKindVersionName(schema).toLowerCase().includes(changes.search.toLowerCase())
      )
      .sort(this.groupKindVersionSorter());
  }

  onSingleSchemaControlChange(changes: ConstraintControlChange): void {
    const properties = this.singleSchema.properties;
    const filteredProperties = Object.keys(properties)
      // filter deleted items
      .filter(propertyKey => {
        if (changes.hideDeleted) {
          return !properties[propertyKey]['x-constraint']?.disabled;
        }
        return true;
      })
      // Filter search
      .filter(propretyKey => propretyKey.toLowerCase().includes(changes.search.toLowerCase()))
      // Collect everything in a new object
      .reduce((previous, propertyKey) => {
        previous[propertyKey] = properties[propertyKey];
        return previous;
      }, {});

    // Create a new schema with the new filtered properties
    this.singleSchemaFiltered = { ...this.singleSchema, properties: filteredProperties };
  }

  private groupKindVersionSorter(): (a: Schema, b: Schema) => number {
    return (a, b) => (this.getGroupKindVersionName(b) > this.getGroupKindVersionName(a) ? -1 : 1);
  }
}
