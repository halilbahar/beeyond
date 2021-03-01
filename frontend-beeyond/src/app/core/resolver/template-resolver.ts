import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { BackendApiService } from '../services/backend-api.service';

@Injectable({
  providedIn: 'root'
})
export class TemplateResolver implements Resolve<Template> {
  constructor(private backendApiService: BackendApiService) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Template> | Template {
    return this.backendApiService.getTemplateById(route.params.id);
  }
}
