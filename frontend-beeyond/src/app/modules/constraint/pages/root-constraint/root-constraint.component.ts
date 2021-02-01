import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';
import { Schema } from 'src/app/shared/models/schema.model';

@Component({
  selector: 'app-root-constraint',
  templateUrl: './root-constraint.component.html',
  styleUrls: ['./root-constraint.component.scss']
})
export class RootConstraintComponent implements OnInit {
  fetching = true;
  schemas: Schema[];
  singleSchema: Schema;

  constructor(
    private validationApiService: ValidationApiService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

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

  navigateFromSchema(schema: Schema) {
    const groupKindVersionNameString = this.getGroupKindVersionName(schema);
    this.navigate(groupKindVersionNameString);
  }

  navigateFromProperty(property: string) {
    this.navigate(property);
  }

  private navigate(path: string) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate([path], { relativeTo: this.route });
  }
}
