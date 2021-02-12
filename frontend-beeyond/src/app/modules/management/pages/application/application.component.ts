import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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
    this.applications = this.route.snapshot.data.applications.sort(
      (a1, a2) => a1.createdAt > a2.createdAt
    );
    this.applicationDataSource = new MatTableDataSource(this.applications);
    this.filterForm = this.fb.group({
      username: [''],
      status: [ApplicationStatus.PENDING],
      fromDate: [null],
      toDate: [null]
    });

    this.availableUsername = this.applications
      .map(application => application.owner.name)
      .filter((name, index, self) => self.indexOf(name) === index);

    this.filterForm.valueChanges.subscribe(() => this.update());

    this.update();
  }

  private update(): void {
    this.selectedRow = null;
    const form: { username: string; status: ApplicationStatus; fromDate: Date; toDate: Date } = this
      .filterForm.value;
    this.applicationDataSource.data = this.applications.filter(({ status, owner, createdAt }) => {
      const nameFilter = form.username ? owner.name.includes(form.username) : true;
      const statusFilter = form.status === ApplicationStatus.ALL || status === form.status;
      const date = new Date(createdAt);

      let fromDateFilter = false;
      if (form.fromDate != null) {
        if (date.getTime() >= form.fromDate.getTime()) {
          fromDateFilter = true;
        }
      } else {
        fromDateFilter = true;
      }

      let toDateFilter = false;
      if (form.toDate != null) {
        if (date.getTime() <= form.toDate.getTime() + 86400000 - 1) {
          toDateFilter = true;
        }
      } else {
        toDateFilter = true;
      }

      const dateFilter = fromDateFilter && toDateFilter;

      return nameFilter && statusFilter && dateFilter;
    });
  }
}
