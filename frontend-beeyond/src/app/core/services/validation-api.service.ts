import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Constraint } from 'src/app/shared/models/constraint.model';
import { Schema } from 'src/app/shared/models/schema.model';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ValidationApiService {
  constructor(private http: HttpClient) {}

  getConstraintForPath(path: string): Observable<Schema[] | Schema> {
    if (path !== '') {
      path = `/${path}`;
    }

    return this.http.get<Schema[] | Schema>(`${environment.validationApiUrl}/constraints` + path);
  }

  createConstraint(path: string, constraint: Constraint): Observable<any> {
    return this.http.post(`${environment.validationApiUrl}/constraints/` + path, constraint);
  }
}
