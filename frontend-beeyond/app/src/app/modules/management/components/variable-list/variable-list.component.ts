import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { ApplicationRange } from 'src/app/shared/models/application-range.model';
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

  fieldData: { value: string; label: string; wildcard: string; description: string }[] = [];

  private template: Template;

  constructor(private backendApiService: BackendApiService, private dialog: MatDialog) {}

  ngOnInit(): void {
    this.backendApiService.getTemplateById(this.templateId).subscribe(template => {
      this.template = template;
      for (const fieldValue of this.fieldValues) {
        const { label, wildcard, description } = template.fields.find(
          aTemplate => aTemplate.id === fieldValue.fieldId
        );
        this.fieldData.push({
          value: fieldValue.value,
          label,
          wildcard,
          description
        });
      }
    });
  }

  openDialog() {
    const templateContent = this.template.content;
    const lines = templateContent.split('\n');

    let content = '';
    const ranges: ApplicationRange[] = [];
    const wildcardRegex = /%(.+?)%/g;

    for (let i = 0; i < lines.length; i++) {
      let line = lines[i];
      let match: RegExpExecArray;

      while ((match = wildcardRegex.exec(line)) !== null) {
        const { wildcard, label, value, description } = this.fieldData.find(
          data => data.wildcard === match[0].replace(/%/g, '')
        );
        line = line.replace(`%${wildcard}%`, value);

        ranges.push({
          lineNumber: i + 1,
          startColumn: match.index + 1,
          endColumn: match.index + 1 + value.length,
          wildcard,
          label,
          description
        });
      }

      content += line + '\n';
    }
    // Remove \n
    content = content.substring(0, content.length - 1);

    this.dialog.open(ApplicationPreviewDialogComponent, {
      data: { content, ranges },
      width: '100%',
      height: '80%',
      autoFocus: false
    });
  }
}
