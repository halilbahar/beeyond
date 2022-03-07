import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Template } from '../../../../shared/models/template.model';
import { ApplicationRange } from '../../../../shared/models/application-range.model';

declare function constrainedEditor(editor: any): any;

@Component({
  selector: 'app-application-review',
  templateUrl: './application-review.component.html',
  styleUrls: ['./application-review.component.scss']
})
export class ApplicationReviewComponent implements OnInit {
  customApplication: CustomApplication | null;
  templateApplication: TemplateApplication | null;

  isPending = false;
  isRunning = false;
  isDenied = false;
  isManagement: boolean;
  redirectPath: string[];

  template: Template;

  fieldData: { value: string; label: string; wildcard: string; description: string }[] = [];

  monacoEditorOptions = {
    language: 'yaml',
    scrollBeyondLastLine: false,
    readOnly: true,
    automaticLayout: true
  };

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private backendApiService: BackendApiService,
    private snackBar: MatSnackBar
  ) {}

  isReadOnly() {
    if (this.isManagement || !this.customApplication) {
      this.monacoEditorOptions.readOnly = true;
      return;
    }
    if (
      this.customApplication.status === ApplicationStatus.DENIED ||
      this.customApplication.status === ApplicationStatus.PENDING
    ) {
      this.monacoEditorOptions.readOnly = false;
      return;
    }
    this.monacoEditorOptions.readOnly = true;
  }

  ngOnInit(): void {
    this.isManagement = this.route.snapshot.data.isManagement;
    this.redirectPath = this.route.snapshot.data.redirectPath;
    const application: CustomApplication | TemplateApplication = this.route.snapshot.data
      .application;
    this.isPending = application.status === ApplicationStatus.PENDING;
    this.isRunning = application.status === ApplicationStatus.RUNNING;
    this.isDenied = application.status === ApplicationStatus.DENIED;

    if ('templateId' in application) {
      this.templateApplication = application;
    } else {
      this.customApplication = application;
    }

    if (this.templateApplication) {
      this.backendApiService
        .getTemplateById(this.templateApplication.templateId)
        .subscribe(template => {
          this.template = template;

          for (const fieldValue of this.templateApplication.fieldValues) {
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

          const templateContent = this.template.content;
          const lines = templateContent.split('\n');

          let content = '';
          const ranges: ApplicationRange[] = [];
          const wildcardRegex = /%(.+?)%/g;

          for (let i = 0; i < lines.length; i++) {
            let line = lines[i];
            let match: RegExpExecArray;

            while ((match = wildcardRegex.exec(line)) != null) {
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

          content = content.substring(0, content.length - 1);
          this.template.content = content;
        });
    }

    this.isReadOnly();
  }

  onEditorInit(editor: any) {
    if (!this.isManagement && this.customApplication.status === ApplicationStatus.STOPPED) {
      const constrainedInstance = constrainedEditor(monaco);
      constrainedInstance.initializeIn(editor);
      const model = editor.getModel();
      const temp = this.customApplication.content.split('\n').map(s => s.trim());
      constrainedInstance.addRestrictionsTo(
        model,
        temp
          .filter(s => s.startsWith('image:'))
          .map(s => ({
            range: [
              temp.indexOf(s) + 1,
              this.customApplication.content
                .split('\n')
                [temp.indexOf(s)].indexOf(s.replace('image: ', '')) + 1,
              temp.indexOf(s) + 1,
              this.customApplication.content
                .split('\n')
                [temp.indexOf(s)].indexOf(s.replace('image: ', '')) +
                s.replace('image: ', '').length +
                1
            ],
            allowMultiline: false
          }))
      );
      model.toggleHighlightOfEditableAreas();
    }
  }

  deny(): void {
    this.backendApiService.denyApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(this.redirectPath);
    });
  }

  approve(): void {
    this.backendApiService.approveApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(this.redirectPath);
    });
  }

  finish(): void {
    this.backendApiService.finishApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(this.redirectPath);
    });
  }

  request(): void {
    this.backendApiService.requestApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(['/profile']).then(navigated => {
        if (navigated) {
          this.snackBar.open(
            'Your application was sent will be reviewed as soon as possible',
            'close',
            { duration: undefined }
          );
        }
      });
    });
  }

  public get application(): CustomApplication | TemplateApplication {
    return this.customApplication || this.templateApplication;
  }
}
