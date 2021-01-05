import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatTableDataSource } from '@angular/material/table';
import { Application } from 'src/app/shared/models/application.model';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';

@Component({
  selector: 'app-application-review',
  templateUrl: './application.component.html',
  styleUrls: ['./application.component.scss']
})
export class ApplicationComponent implements OnInit {
  applications: Application[];
  applicationDataSource: MatTableDataSource<Application>;
  columnsToDisplay = ['id', 'owner', 'status'];

  filterForm: FormGroup;
  availableUsername: string[];
  statuses: ApplicationStatus[] = [
    ApplicationStatus.ALL,
    ApplicationStatus.PENDING,
    ApplicationStatus.DENIED,
    ApplicationStatus.APPROVED
  ];

  selectedRow: number | null;

  constructor(private route: ActivatedRoute, private fb: FormBuilder) {}

  ngOnInit(): void {
    this.applications = this.route.snapshot.data.applications;
    this.applicationDataSource = new MatTableDataSource(this.applications);
    this.filterForm = this.fb.group({
      username: [''],
      status: [ApplicationStatus.ALL]
    });

    this.availableUsername = this.applications
      .map(application => application.owner.name)
      .filter((name, index, self) => self.indexOf(name) === index);

    this.filterForm.valueChanges.subscribe(
      (form: { username: string; status: ApplicationStatus }) => {
        this.selectedRow = null;
        this.applicationDataSource.data = this.applications.filter(({ status, owner }) => {
          const nameFilter = form.username ? owner.name.includes(form.username) : true;
          const statusFilter = form.status === ApplicationStatus.ALL || status === form.status;

          return nameFilter && statusFilter;
        });
      }
    );
  }
}
