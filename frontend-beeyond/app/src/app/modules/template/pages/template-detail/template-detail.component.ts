import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { ThemeService } from '../../../../core/services/theme.service';

@Component({
  selector: 'app-template-detail',
  templateUrl: './template-detail.component.html',
  styleUrls: ['./template-detail.component.scss']
})
export class TemplateDetailComponent implements OnInit {
  template: Template;

  monacoEditorOptions = {
    language: 'yaml',
    scrollBeyondLastLine: false,
    readOnly: true,
    theme: this.themeService.isDarkTheme.value ? 'vs-dark' : 'vs-light'
  };

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private backendApiService: BackendApiService,
    private themeService: ThemeService
  ) {}

  ngOnInit(): void {
    this.themeService.isDarkTheme.subscribe(value => {
      this.monacoEditorOptions = {
        ...this.monacoEditorOptions,
        theme: value ? 'vs-dark' : 'vs-light'
      };
    });
    this.template = this.route.snapshot.data.template;
  }

  delete() {
    this.backendApiService.deleteTemplate(this.template.id).subscribe(() => {
      this.router.navigate(['/template']);
    });
  }
}
