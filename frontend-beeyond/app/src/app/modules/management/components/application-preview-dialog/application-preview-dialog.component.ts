import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ApplicationRange } from 'src/app/shared/models/application-range.model';

declare let monaco: any;

@Component({
  selector: 'app-application-preview-dialog',
  templateUrl: './application-preview-dialog.component.html',
  styleUrls: ['./application-preview-dialog.component.scss']
})
export class ApplicationPreviewDialogComponent {
  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  constructor(
    @Inject(MAT_DIALOG_DATA)
    public data: {
      content: string;
      ranges: ApplicationRange[];
    }
  ) {}

  onEditorInit(editor: any) {
    const decorations = this.data.ranges.map(applicationRange => ({
      range: new monaco.Range(
        applicationRange.lineNumber,
        applicationRange.startColumn,
        applicationRange.lineNumber,
        applicationRange.endColumn
      ),
      options: {
        linesDecorationsClassName: 'monaco-application-line-decoration',
        inlineClassName: 'monaco-application-inline',
        hoverMessage: {
          value:
            '```yaml\n' +
            `   wildcard: ${applicationRange.wildcard}\n` +
            `      label: ${applicationRange.label}\n` +
            `description: ${applicationRange.description}` +
            '\n```'
        }
      }
    }));

    editor.deltaDecorations([], decorations);
  }
}
