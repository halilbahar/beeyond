import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-blueprint-template',
  templateUrl: './blueprint-template.component.html',
  styleUrls: ['./blueprint-template.component.css']
})
export class BlueprintTemplateComponent implements OnInit {

  template = '';

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.url.subscribe(url => this.template = url[1].path);
  }
}
