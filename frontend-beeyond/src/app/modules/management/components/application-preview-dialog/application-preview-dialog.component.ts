import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-application-preview-dialog',
  templateUrl: './application-preview-dialog.component.html',
  styleUrls: ['./application-preview-dialog.component.scss']
})
export class ApplicationPreviewDialogComponent {
  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  constructor(@Inject(MAT_DIALOG_DATA) public content: string) {}
}
