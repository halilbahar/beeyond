import { Component, Input, OnInit } from '@angular/core';
import { Application } from '../../../../shared/models/application.model';

@Component({
  selector: 'app-application-attributes',
  templateUrl: './application-attributes.component.html',
  styleUrls: ['./application-attributes.component.scss']
})
export class ApplicationAttributesComponent implements OnInit {
  @Input() public application: Application;

  constructor() {}

  ngOnInit(): void {}
}
