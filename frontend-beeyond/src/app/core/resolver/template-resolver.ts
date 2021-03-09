import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { BackendApiService } from '../services/backend-api.service';
import { tap } from 'rxjs/operators';
import { ProgressBarService } from '../services/progress-bar.service';

@Injectable({
  providedIn: 'root'
})
export class TemplateResolver implements Resolve<Template> {
  constructor(
    private backendApiService: BackendApiService,
    private progressBarService: ProgressBarService
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Template> | Template {
    this.progressBarService.start();
    return this.backendApiService
      .getTemplateById(route.params.id)
      .pipe(tap(() => this.progressBarService.finish()));
  }
}
