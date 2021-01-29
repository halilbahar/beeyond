import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { ApiService } from '../../../../core/services/api.service';

@Component({
  selector: 'app-template-detail',
  templateUrl: './template-detail.component.html',
  styleUrls: ['./template-detail.component.scss']
})
export class TemplateDetailComponent implements OnInit {
  template: Template;

  monacoEditorOptions = { language: 'yaml', scrollBeyondLastLine: false, readOnly: true };

  constructor(private router: Router, private route: ActivatedRoute, private service: ApiService) {}

  ngOnInit(): void {
    this.template = this.route.snapshot.data.template;
  }

  delete() {
    this.service.deleteTemplate(this.template.id).subscribe(() => {
      this.router.navigate(['/template']);
    });
    /*this.service.deleteTemplate(this.template.id).subscribe(() => {
      this.router.navigate(['/template']);
    });*/
  }
}
