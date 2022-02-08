import { ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { MatTableDataSource } from '@angular/material/table';
import { ActivatedRoute } from '@angular/router';
import { ApplicationStatus } from 'src/app/shared/models/application-status.enum';
import { Application } from 'src/app/shared/models/application.model';
import { BaseComponent } from '../../../../core/services/base.component';
import { MediaMatcher } from '@angular/cdk/layout';

@Component({
  selector: 'app-application-content',
  templateUrl: './application-content.component.html',
  styleUrls: ['./application-content.component.scss']
})
export class ApplicationContentComponent extends BaseComponent implements OnInit {
  @Input() isAdmin = true;

  applications: Application[];
  applicationDataSource: MatTableDataSource<Application>;
  columnsToDisplay = ['id', 'status', 'startedAt', 'finishedAt'];

  filterForm: FormGroup;
  availableUsername: string[];
  statuses: ApplicationStatus[] = [
    ApplicationStatus.ALL,
    ApplicationStatus.PENDING,
    ApplicationStatus.DENIED,
    ApplicationStatus.RUNNING,
    ApplicationStatus.FINISHED
  ];

  selectedRow: number | null;

  constructor(private route: ActivatedRoute,
              private fb: FormBuilder,
              changeDetectorRef: ChangeDetectorRef,
              media: MediaMatcher) {
    super(changeDetectorRef, media);
    if(!super.mobileQuery?.matches){
      this.columnsToDisplay = ['id', 'status'];
    }
  }

  ngOnInit(): void {
    if (this.isAdmin) {
      this.columnsToDisplay.splice(1, 0, 'owner');
    }

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
    this.applicationDataSource.data = this.applications.filter(({status, owner, createdAt}) => {
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
