import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { ApiService } from '../../../../core/services/api.service';

@Component({
  selector: 'app-template-detail',
  templateUrl: './template-detail.component.html',
  styleUrls: ['./template-detail.component.scss']
})
export class TemplateDetailComponent implements OnInit {
  @Input() routePath = '';

  template: Template;
  id: number;
  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private apiService: ApiService
  ) {}

  ngOnInit(): void {
    this.id = Number(this.route.snapshot.paramMap.get('id'));
    this.apiService.getTemplateById(this.id).subscribe(template => {
      this.template = template;
    });
  }
}
