import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }

  getTemplates(): Observable<Template[]> {
    return this.http.get<Template[]>(`${environment.apiUrl}/template`);
  }

  getTemplateById(id: number): Observable<Template> {
    return this.http.get<Template>(`${environment.apiUrl}/template/${id}`);
  }

  createTemplateApplication(body: object): Observable<null> {
    return this.http.post<null>(`${environment.apiUrl}/application/template`, body);
  }

  createCustomApplication(body: object): Observable<any> {
    return this.http.post<any>(`${environment.apiUrl}/application/custom`, body);
  }

  createTemplate(body: object): Observable<any> {
    return this.http.post<any>(`${environment.apiUrl}/template`, body);
  }
}
