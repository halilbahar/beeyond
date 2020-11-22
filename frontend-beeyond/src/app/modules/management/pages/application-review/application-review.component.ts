import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Application } from '../../../../shared/models/application.model';
import { MatButtonModule } from '@angular/material/button';
import { ApiService } from 'src/app/core/services/api.service';

@Component({
  selector: 'app-application-review',
  templateUrl: './application-review.component.html',
  styleUrls: ['./application-review.component.scss']
})
export class ApplicationReviewComponent implements OnInit {
  application: any;
  template: boolean = false;

  constructor(private route: ActivatedRoute, private service: ApiService) {}

  ngOnInit(): void {
    this.application = this.route.snapshot.data.application;
    if (this.application.templateId != null) {
      this.template = true;
    }
  }

  deny(): void {
    this.service.denyApplicationById(this.application.id).subscribe(console.log);
  }

  approve(): void {
    this.service.approveApplicationById(this.application.id).subscribe(console.log);
  }
}
