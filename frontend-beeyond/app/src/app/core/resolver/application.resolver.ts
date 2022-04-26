import { Injectable } from '@angular/core';
import { BackendApiService } from '../services/backend-api.service';
import { ActivatedRouteSnapshot, Resolve } from '@angular/router';
import { Application } from 'src/app/shared/models/application.model';
import { Observable } from 'rxjs';
import { ProgressBarService } from '../services/progress-bar.service';
import { tap } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class ApplicationResolver implements Resolve<Application[]> {
  constructor(
    private backendApiService: BackendApiService,
    private progressBarService: ProgressBarService
  ) {}

  resolve(route: ActivatedRouteSnapshot): Observable<Application[]> {
    this.progressBarService.start();
    return this.backendApiService
      .getApplications(route.data.all)
      .pipe(tap(() => this.progressBarService.finish()));
  }
}
