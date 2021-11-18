import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';
import { Application } from 'src/app/shared/models/application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';

@Component({
  selector: 'app-application-review',
  templateUrl: './application-review.component.html',
  styleUrls: ['./application-review.component.scss']
})
export class ApplicationReviewComponent implements OnInit {
  customApplication: CustomApplication | null;
  templateApplication: TemplateApplication | null;

  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  isPending = false;
  isRunning = false;
  isManagement: boolean;
  redirectPath: string[];

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private backendApiService: BackendApiService
  ) {}

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
    this.backendApiService.stopApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(this.redirectPath);
    });
  }

  private get application(): CustomApplication | TemplateApplication {
    return this.customApplication || this.templateApplication;
  }
}
