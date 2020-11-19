import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Application } from '../../../../shared/models/application.model';
import { ApiService } from 'src/app/core/services/api.service';

@Component({
  selector: 'app-application-review',
  templateUrl: './application.component.html',
  styleUrls: ['./application.component.scss']
})
export class ApplicationComponent implements OnInit {
  applications: any;
  columnsToDisplay = ['id', 'owner.name', 'status'];

  constructor(private route: ActivatedRoute, private router:Router, private service: ApiService) {}

  ngOnInit(): void {
    this.applications = this.route.snapshot.data.applications;
  }

  routeTo(id: number) {
    this.router.navigate(["/application/review/"+id]).then(console.log);;
  }
}
