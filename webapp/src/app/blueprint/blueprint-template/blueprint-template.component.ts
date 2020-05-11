import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { HttpApiService } from '../../service/http-api.service';
import { BlueprintTemplate } from './blueprint-template.model';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-blueprint-template',
  templateUrl: './blueprint-template.component.html',
  styleUrls: ['./blueprint-template.component.css']
})
export class BlueprintTemplateComponent implements OnInit {

  template = '';
  form = {name: '', replica: 1};

  constructor(
    private route: ActivatedRoute,
    private httpApiService: HttpApiService,
    private snackBar: MatSnackBar
  ) { }

  ngOnInit(): void {
    this.route.url.subscribe(url => this.template = url[1].path);
  }

  submit() {
    this.httpApiService.postTemplate(new BlueprintTemplate(this.form.name, this.form.replica, this.template))
      .subscribe(value => {
        this.snackBar.open('Application was sent!', '', {duration: 2000});
      });
  }
}
