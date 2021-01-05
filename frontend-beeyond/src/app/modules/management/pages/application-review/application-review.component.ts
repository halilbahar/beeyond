import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from 'src/app/core/services/api.service';
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

  constructor(private route: ActivatedRoute, private router: Router, private service: ApiService) {}

  ngOnInit(): void {
    const application = this.route.snapshot.data.application;
    if ('templateId' in application) {
      this.templateApplication = application;
    } else {
      this.customApplication = application;
    }
  }

  deny(): void {
    this.service.denyApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(['/management']);
    });
  }

  approve(): void {
    this.service.approveApplicationById(this.application.id).subscribe(() => {
      this.router.navigate(['/management']);
    });
  }

  private get application(): CustomApplication | TemplateApplication {
    return this.customApplication || this.templateApplication;
  }
}
