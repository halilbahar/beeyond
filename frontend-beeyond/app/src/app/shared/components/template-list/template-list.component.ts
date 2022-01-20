import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
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
  @Output() templateIdEvent = new EventEmitter<number>();
  selectedTemplateId: number;

  constructor(private router: Router) {}

  ngOnInit(): void {}

  route(id: number) {
    this.templateIdEvent.emit(id);
    this.selectedTemplateId = id;
    if (!this.routingEnabled) {
      return;
    }
    this.router.navigate([this.routePath.replace('{id}', String(id))]).then(console.log);
  }
}
