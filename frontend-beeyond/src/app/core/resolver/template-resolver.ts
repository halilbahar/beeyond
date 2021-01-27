import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';

@Injectable({
  providedIn: 'root'
})
export class TemplateResolver implements Resolve<Template> {
  constructor(private service: ApiService) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Template> | Template {
    return this.service.getTemplateById(route.params.id);
  }
}
