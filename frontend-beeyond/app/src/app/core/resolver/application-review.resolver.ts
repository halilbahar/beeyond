import { Injectable } from '@angular/core';
import { BackendApiService } from '../services/backend-api.service';
import { Resolve, ActivatedRouteSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { tap } from 'rxjs/operators';
import { ProgressBarService } from '../services/progress-bar.service';

@Injectable({ providedIn: 'root' })
export class ApplicationReviewResolver implements Resolve<TemplateApplication | CustomApplication> {
  constructor(
    private backendApiService: BackendApiService,
    private progressBarService: ProgressBarService
  ) {}

  resolve(route: ActivatedRouteSnapshot): Observable<TemplateApplication | CustomApplication> {
    this.progressBarService.start();
    return this.backendApiService
      .getApplicationById(route.params.id)
      .pipe(tap(() => this.progressBarService.finish()));
  }
}
