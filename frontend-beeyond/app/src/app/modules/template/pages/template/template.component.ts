import { Component, OnInit } from '@angular/core';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Template } from '../../../../shared/models/template.model';

@Component({
  selector: 'app-template',
  templateUrl: './template.component.html',
  styleUrls: ['./template.component.scss']
})
export class TemplateComponent implements OnInit {
  templates: Template[] = [];

  constructor(private backendApiService: BackendApiService) {}

  ngOnInit(): void {
    this.backendApiService
      .getTemplates()
      .subscribe(templates => (this.templates = templates.filter(template => !template.deleted)));
  }
}
