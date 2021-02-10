import { Injectable } from '@angular/core';
import { BackendApiService } from '../services/backend-api.service';
import { Resolve, ActivatedRouteSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';

@Injectable({ providedIn: 'root' })
export class ApplicationReviewResolver implements Resolve<TemplateApplication | CustomApplication> {
  constructor(private backendApiService: BackendApiService) {}

  resolve(route: ActivatedRouteSnapshot): Observable<TemplateApplication | CustomApplication> {
    return this.backendApiService.getApplicationById(route.params.id);
  }
}
