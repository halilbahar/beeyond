import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ApiService } from 'src/app/core/services/api.service';
import { TemplateFieldValue } from 'src/app/shared/models/template-field-value.model';
import { Template } from 'src/app/shared/models/template.model';
import { ApplicationPreviewDialogComponent } from '../application-preview-dialog/application-preview-dialog.component';

@Component({
  selector: 'app-variable-list',
  templateUrl: './variable-list.component.html',
  styleUrls: ['./variable-list.component.scss']
})
export class VariableListComponent implements OnInit {
  @Input() private fieldValues: TemplateFieldValue[];
  @Input() private templateId: number;

  fieldData: { value: string; label: string; wildcard: string }[] = [];

  private template: Template;

  constructor(private service: ApiService, private dialog: MatDialog) {}

  ngOnInit(): void {
    this.service.getTemplateById(this.templateId).subscribe(template => {
      this.template = template;
      for (const fieldValue of this.fieldValues) {
        const { label, wildcard } = template.fields.find(
          aTemplate => aTemplate.id === fieldValue.fieldId
        );
        this.fieldData.push({
          value: fieldValue.value,
          label,
          wildcard
        });
      }
    });
  }

  openDialog() {
    let content = this.template.content;
    for (const data of this.fieldData) {
      content = content.replace(`%${data.wildcard}%`, data.value);
    }

    this.dialog.open(ApplicationPreviewDialogComponent, {
      data: content,
      width: '100%',
      height: '80%',
      autoFocus: false
    });
  }
}
