import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';
import { Schema } from 'src/app/shared/models/schema.model';

@Component({
  selector: 'app-constraint',
  templateUrl: './constraint.component.html',
  styleUrls: ['./constraint.component.scss']
})
export class ConstraintComponent implements OnInit {
  fetching = true;
  schemas: Schema[];
  singleSchema: Schema;

  constructor(private validationApiService: ValidationApiService, private route: ActivatedRoute) {}

  ngOnInit(): void {
    const segments = this.route.snapshot.url.map(segment => segment.path);
    const constraintPath = segments.join('/');

    this.validationApiService.getConstraintForPath(constraintPath).subscribe(schemas => {
      this.fetching = false;
      if (schemas instanceof Array) {
        this.schemas = schemas;
      } else {
        this.singleSchema = schemas;
      }
    });
  }

  get propertyKeys() {
    return Object.keys(this.singleSchema.properties);
  }

  getGroupKindVersionName(schema: Schema) {
    const { group, kind, version } = schema['x-kubernetes-group-version-kind'][0];

    let groupString = '';
    if (group !== '') {
      groupString = '-' + group;
    }

    return kind + groupString + '-' + version;
  }
}
