import { Component, Input, OnInit } from '@angular/core';
import { Template } from '../../models/template.model';

@Component({
  selector: 'app-template-list',
  templateUrl: './template-list.component.html',
  styleUrls: ['./template-list.component.scss']
})
export class TemplateListComponent implements OnInit {

  @Input() templates: Template[];

  constructor() { }

  ngOnInit(): void {
  }
}
