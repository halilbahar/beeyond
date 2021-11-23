import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Constraint } from 'src/app/shared/models/constraint.model';
import { Schema } from 'src/app/shared/models/schema.model';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root'
})
export class ValidationApiService {
  private readonly validationApiUrl: string;

  constructor(private http: HttpClient, private configService: ConfigService) {
    this.validationApiUrl = configService.config.validationApiUrl;
  }

  getConstraintForPath(path: string): Observable<Schema[] | Schema> {
    if (path !== '') {
      path = `/${path}`;
    }

    return this.http.get<Schema[] | Schema>(`${this.validationApiUrl}/constraints` + path);
  }

  createConstraint(path: string, constraint: Constraint): Observable<void> {
    return this.http.post<void>(`${this.validationApiUrl}/constraints/` + path, constraint);
  }

  deleteConstraint(path: string): Observable<void> {
    return this.http.delete<void>(`${this.validationApiUrl}/constraints/` + path);
  }

  toggleConstraint(path: string): Observable<void> {
    return this.http.patch<void>(`${this.validationApiUrl}/constraints/` + path, null);
  }
}
