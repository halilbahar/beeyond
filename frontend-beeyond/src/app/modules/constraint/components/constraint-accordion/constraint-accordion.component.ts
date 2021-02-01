import { Component, Input, OnInit } from '@angular/core';
import { Schema } from 'src/app/shared/models/schema.model';

@Component({
  selector: 'app-constraint-accordion',
  templateUrl: './constraint-accordion.component.html',
  styleUrls: ['./constraint-accordion.component.scss']
})
export class ConstraintAccordionComponent implements OnInit {
  @Input() schemas: Schema[];

  constructor() {}

  ngOnInit(): void {}

  getGroupKindVersionName(schema: Schema) {
    const { group, kind, version } = schema['x-kubernetes-group-version-kind'][0];

    let groupString = '';
    if (group !== '') {
      groupString = '-' + group;
    }

    return kind + groupString + '-' + version;
  }
}
