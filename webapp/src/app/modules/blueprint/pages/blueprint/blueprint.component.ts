import { Component, OnInit } from '@angular/core';
import { Template } from '../../../../shared/models/template.model';
import { ApiService } from '../../../../core/services/api.service';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.scss']
})
export class BlueprintComponent implements OnInit {

  templates: Template[] = [];
  customTemplate = {
    content: '',
    note: ''
  };

  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    this.apiService.getTemplates().subscribe(templates => this.templates = templates);
  }

  sendCustomTemplate(): void {
  }
}
