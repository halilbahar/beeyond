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
  @Input() fieldValues: TemplateFieldValue[];
  @Input() templateId: number;

  fieldData: { value: string; fieldName: string }[] = [];

  constructor(private service: ApiService) {}

  ngOnInit(): void {
    this.service.getTemplateById(this.templateId).subscribe(template => {
      for (const fieldValue of this.fieldValues) {
        const templateField = template.fields.find(aTemplate => aTemplate.id === fieldValue.fieldId);
        this.fieldData.push({
          value: fieldValue.value,
          fieldName: templateField.label
        });
      }
    });
  }
}
