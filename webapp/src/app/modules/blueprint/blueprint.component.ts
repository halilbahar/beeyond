import { Component, OnInit } from '@angular/core';
import { HttpApiService } from '../../service/http-api.service';
import { Template } from './template.model';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.css']
})
export class BlueprintComponent implements OnInit {

  templates: Template[] = [];

  constructor(private httpApiService: HttpApiService) { }

  ngOnInit(): void {
    this.httpApiService.getAllTemplates().subscribe(value => this.templates = value);
  }
}
