import { Component, OnInit } from '@angular/core';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';
import { Schema } from 'src/app/shared/models/schema.model';

@Component({
  selector: 'app-root-constraint',
  templateUrl: './root-constraint.component.html',
  styleUrls: ['./root-constraint.component.scss']
})
export class RootConstraintComponent implements OnInit {
  schemas: Schema[];

  constructor(private validationApiService: ValidationApiService) {}

  ngOnInit(): void {
    this.validationApiService.getRootConstraints().subscribe(schemas => (this.schemas = schemas));
  }
}
