import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';

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
  isManagement: boolean;
  redirectPath: string[];

  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private backendApiService: BackendApiService
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

    if ('templateId' in application) {
      this.templateApplication = application;
    } else {
      this.customApplication = application;
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

  private get application(): CustomApplication | TemplateApplication {
    return this.customApplication || this.templateApplication;
  }
}
