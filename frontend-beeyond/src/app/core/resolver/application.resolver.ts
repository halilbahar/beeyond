import { Injectable } from '@angular/core';
import { BackendApiService } from '../services/backend-api.service';
import { Resolve } from '@angular/router';
import { Application } from 'src/app/shared/models/application.model';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class ApplicationResolver implements Resolve<Application[]> {
  constructor(private backendApiService: BackendApiService) {}

  resolve(): Observable<Application[]> {
    return this.backendApiService.getApplications();
  }
}
