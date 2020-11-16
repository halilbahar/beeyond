import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Resolve, ActivatedRouteSnapshot } from '@angular/router';
import { Application } from 'src/app/shared/models/application.model';
import { Observable } from 'rxjs';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';

@Injectable({ providedIn: 'root' })
export class ApplicationReviewResolver implements Resolve<TemplateApplication | CustomApplication> {
  constructor(
    private service: ApiService
  ) { }

  resolve(route: ActivatedRouteSnapshot): Observable<TemplateApplication | CustomApplication> {
    return this.service.getApplicationById(route.params['id']);
  }
}
