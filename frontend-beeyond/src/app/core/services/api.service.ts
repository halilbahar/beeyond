import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { Application } from 'src/app/shared/models/application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {}

  getTemplates(): Observable<Template[]> {
    return this.http.get<Template[]>(`${environment.apiUrl}/template`);
  }

  getTemplateById(id: number): Observable<Template> {
    return this.http.get<Template>(`${environment.apiUrl}/template/${id}`);
  }

  // TODO: type body
  createTemplateApplication(body: Record<string, unknown>): Observable<null> {
    return this.http.post<null>(`${environment.apiUrl}/application/template`, body);
  }

  // TODO: type body
  createCustomApplication(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${environment.apiUrl}/application/custom`, body);
  }

  // TODO: type body
  createTemplate(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${environment.apiUrl}/template`, body);
  }

  deleteTemplate(id: number): Observable<any> {
    return this.http.delete<any>(`${environment.apiUrl}/template/${id}`, null);
  }

  getApplications(): Observable<Application[]> {
    return this.http.get<Application[]>(`${environment.apiUrl}/application`);
  }

  getApplicationById(id: number): Observable<TemplateApplication | CustomApplication> {
    return this.http.get<TemplateApplication | CustomApplication>(
      `${environment.apiUrl}/application/${id}`
    );
  }

  approveApplicationById(id: number): Observable<null> {
    return this.http.patch<any>(`${environment.apiUrl}/application/approve/${id}`, null);
  }

  denyApplicationById(id: number): Observable<null> {
    return this.http.patch<any>(`${environment.apiUrl}/application/deny/${id}`, null);
  }
}
