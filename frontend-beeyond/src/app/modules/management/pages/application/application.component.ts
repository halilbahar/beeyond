import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from 'src/app/core/services/api.service';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { MatTableDataSource } from '@angular/material/table';
import { Observable } from 'rxjs';
import { map, startWith } from 'rxjs/operators';
import { MatSort } from '@angular/material/sort';

@Component({
  selector: 'app-application-review',
  templateUrl: './application.component.html',
  styleUrls: ['./application.component.scss']
})
export class ApplicationComponent implements OnInit {
  filterdIdOptions: Observable<number[]>;
  filterdUserOptions: Observable<string[]>;
  statuses: string[] = ['ALL', 'PENDING', 'DENIED', 'APPROVED'];
  columnsToDisplay = ['id', 'owner.name', 'status'];
  applications: any;
  user: string;
  id: string;
  status: string;
  dataSource: any;
  userSearchInput: FormControl = new FormControl();
  searchInput: FormControl = new FormControl();
  @ViewChild(MatSort) sort: MatSort;

  filterForm: FormGroup;

  constructor(private route: ActivatedRoute, private router: Router, private _fb: FormBuilder) { }
  
  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
  }
  ngOnInit(): void {
    this.applications = this.route.snapshot.data.applications;
    this.dataSource = new MatTableDataSource(this.applications);
    this.dataSource.filterPredicate = this.customFiltered();
    this.filterForm = this._fb.group({
      searchInput: [],
      userSearchInput: [],
      statusSearchInput: []
    });
    this.filterdIdOptions = this.searchInput.valueChanges.pipe(
      startWith(''),
      map(value => this.filterId(value))
    );
    this.filterdUserOptions = this.userSearchInput.valueChanges.pipe(
      startWith(''),
      map(value => this.filterUser(value))
    );
  }

  customFiltered() {
    return (data, filter) => {
      if (this.id && this.user && this.status) {
        return String(data.id).includes(this.id) && data.owner.name.includes(this.user) && (data.status == this.status || this.status == 'ALL');
      }else if (this.id && this.status) {
        return String(data.id).includes(this.id) && (data.status == this.status || this.status == 'ALL');
      }else if (this.id && this.user) {
        return String(data.id).includes(this.id) && data.owner.name.includes(this.user);
      }else if (this.user && this.status) {
        return data.owner.name.includes(this.user) && (data.status == this.status || this.status == 'ALL');
      }else if (this.id) {
        return String(data.id).includes(this.id);
      }else if (this.user) {
        return data.owner.name.includes(this.user);
      }else if (this.status) {
        return data.status == this.status || this.status == 'ALL';
      }
      return true;
    }
  }

  applyFilter(id: string, user: string, status: string) {
    this.dataSource.filter = id + "," + user;
  }

  private filterId(value: number): number[] {
    return this.applications.map((application) => application.id).filter(id => String(id).includes(String(value))).sort((a, b) => a - b);
  }

  private filterUser(value: string): string[] {
    return this.applications.map((application) => application.owner.name).filter(user => user.includes(value)).filter((v, i, a) => a.indexOf(v) === i);;
  }

  routeTo(id: number) {
    this.router.navigate(["/management/review/" + id]).then(console.log);;
  }
}
