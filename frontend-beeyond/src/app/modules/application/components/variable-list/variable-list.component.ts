import { Component, Input, OnInit } from '@angular/core';
import { ApiService } from 'src/app/core/services/api.service';
import { TemplateFieldValue } from 'src/app/shared/models/template-field-value.model';
import { TemplateField } from 'src/app/shared/models/template-field.model';

@Component({
  selector: 'app-variable-list',
  templateUrl: './variable-list.component.html',
  styleUrls: ['./variable-list.component.scss']
})
export class VariableListComponent implements OnInit {

  @Input() fields: TemplateFieldValue[];
  @Input() templateId: number;

  data: any;

  fieldNames: TemplateField[];

  constructor(private service: ApiService) { }

  ngOnInit(): void {
    this.data = [];
    this.service.getTemplateById(this.templateId).subscribe(template => {
      this.fieldNames = template.fields;
      this.fields.forEach((elt, i) => {
        this.data.push({ field: elt, fieldName: this.fieldNames[i] });
      });
    });
  }
}
