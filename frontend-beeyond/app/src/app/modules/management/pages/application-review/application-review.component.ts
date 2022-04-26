import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Template } from '../../../../shared/models/template.model';
import { ThemeService } from '../../../../core/services/theme.service';
import { MatDialog } from '@angular/material/dialog';
import { ApplicationDenyDialogComponent } from '../../components/application-deny-dialog/application-deny-dialog.component';
import { Observable } from 'rxjs';
import { FormControl, Validators } from '@angular/forms';
import { ConfigService } from '../../../../core/services/config.service';

declare function constrainedEditor(editor: any): any;

declare let monaco: any;

@Component({
  selector: 'app-application-review',
  templateUrl: './application-review.component.html',
  styleUrls: ['./application-review.component.scss']
})
export class ApplicationReviewComponent implements OnInit {
  customApplication: CustomApplication | null;
  customForm: FormControl;
  templateApplication: TemplateApplication | null;

  isPending = false;
  isRunning = false;
  isDenied = false;
  isManagement: boolean;
  redirectPath: string[];
  message: string;

  running: ApplicationStatus = ApplicationStatus.RUNNING;
  stopped: ApplicationStatus = ApplicationStatus.STOPPED;
  denied: ApplicationStatus = ApplicationStatus.DENIED;

  template: Template;

  fieldData: { value: string; label: string; wildcard: string; description: string }[] = [];

  monacoEditorOptions = {
    language: 'yaml',
    scrollBeyondLastLine: false,
    readOnly: true,
    automaticLayout: true,
    theme: this.themeService.isDarkTheme.value ? 'vs-dark' : 'vs-light'
  };

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private backendApiService: BackendApiService,
    private snackBar: MatSnackBar,
    private themeService: ThemeService,
    public configService: ConfigService,
    public dialog: MatDialog
  ) {
    this.themeService.isDarkTheme.subscribe(value => {
      this.monacoEditorOptions = {
        ...this.monacoEditorOptions,
        theme: value ? 'vs-dark' : 'vs-light'
      };
    });
  }

  public get application(): CustomApplication | TemplateApplication {
    return this.customApplication || this.templateApplication;
  }

  isReadOnly() {
    if (this.isManagement || !this.customApplication) {
      this.monacoEditorOptions.readOnly = true;
      return;
    }
    if (
      this.customApplication.status === ApplicationStatus.DENIED ||
      this.customApplication.status === ApplicationStatus.STOPPED ||
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
    const application: CustomApplication | TemplateApplication =
      this.route.snapshot.data.application;
    this.isPending = application.status === ApplicationStatus.PENDING;
    this.isRunning = application.status === ApplicationStatus.RUNNING;
    this.isDenied = application.status === ApplicationStatus.DENIED;

    if ('templateId' in application) {
      this.templateApplication = application;
    } else {
      this.customApplication = application;
      this.customForm = new FormControl(this.customApplication.content, Validators.required);
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
    const dialogRef = this.dialog.open(ApplicationDenyDialogComponent, {
      width: '600px',
      height: '500px',
      data: { message: this.message }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result !== undefined) {
        this.backendApiService.denyApplicationById(this.application.id, result).subscribe(() => {
          this.router.navigate(this.redirectPath);
        });
      }
    });
  }

  approve(): void {
    this.backendApiService.approveApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(this.redirectPath);
    });
  }

  request(): void {
    this.save().subscribe(() =>
      this.backendApiService.requestApplicationById(this.application.id).subscribe(() => {
        this.router.navigate(['/profile']).then(navigated => {
          if (navigated) {
            this.snackBar.open(
              'Your application was sent will be reviewed as soon as possible',
              'close',
              {
                duration: 2000,
                panelClass: ['mat-drawer-container']
              }
            );
          }
        });
      })
    );
  }

  stop(): void {
    this.backendApiService.stopApplicationById(this.application.id).subscribe(() => {
      const currentUrl = this.router.url;
      this.router.navigateByUrl('/', { skipLocationChange: true }).then(() => {
        this.router.navigate([currentUrl]);
      });
    });
  }

  finish(): void {
    //TODO Add dialog question, do you really wanna finish? SlEm
    this.backendApiService.finishApplicationById(this.application.id).subscribe(() => {
      const currentUrl = this.router.url;
      this.router.navigateByUrl('/', { skipLocationChange: true }).then(() => {
        this.router.navigate([currentUrl]);
      });
    });
  }

  start(): void {
    this.backendApiService.startApplicationById(this.application.id).subscribe(() => {
      const currentUrl = this.router.url;
      this.router.navigateByUrl('/', { skipLocationChange: true }).then(() => {
        this.router.navigate([currentUrl]);
      });
    });
  }

  save(): Observable<void> {
    return this.backendApiService.saveApplicationById(this.application.id, {
      content: this.customForm.value
    });
  }
}
