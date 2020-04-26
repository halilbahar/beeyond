import { Component, OnInit, ViewChild } from '@angular/core';
import { MatSelectionList } from '@angular/material/list';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.css']
})
export class BlueprintComponent implements OnInit {

  @ViewChild('blueprints') blueprints: MatSelectionList;

  constructor() { }

  ngOnInit(): void {
  }

}
