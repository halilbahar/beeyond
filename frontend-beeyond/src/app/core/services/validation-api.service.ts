import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Schema } from 'src/app/shared/models/schema.model';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ValidationApiService {
  constructor(private http: HttpClient) {}

  getRootConstraints(): Observable<Schema[]> {
    return this.http.get<Schema[]>(`${environment.validationApiUrl}/constraints`);
  }
}
