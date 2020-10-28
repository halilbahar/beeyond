import { Component, Input, OnInit } from '@angular/core';
import { Template } from '../../models/template.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-template-list',
  templateUrl: './template-list.component.html',
  styleUrls: ['./template-list.component.scss']
})
export class TemplateListComponent implements OnInit {

  @Input() templates: Template[];
  @Input() routingEnabled = false;
  @Input() routePath = '';

  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  route(id: number) {
    if (!this.routingEnabled) {
      return;
    }
    this.router.navigate([this.routePath.replace('{id}', String(id))])
      .then(console.log);
  }
}
