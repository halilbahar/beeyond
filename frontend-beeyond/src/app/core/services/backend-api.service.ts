import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { Application } from 'src/app/shared/models/application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { Namespace } from 'src/app/shared/models/namespace.model';
import { User } from 'src/app/shared/models/user.model';
import { NamespaceSave } from 'src/app/shared/models/namespace-save.model';

@Injectable({
  providedIn: 'root'
})
export class BackendApiService {
  constructor(private http: HttpClient) {}

  getTemplates(): Observable<Template[]> {
    return this.http.get<Template[]>(`${environment.backendApiUrl}/template`);
  }

  getTemplateById(id: number): Observable<Template> {
    return this.http.get<Template>(`${environment.backendApiUrl}/template/${id}`);
  }

  // TODO: type body
  createTemplateApplication(body: Record<string, unknown>): Observable<null> {
    return this.http.post<null>(`${environment.backendApiUrl}/application/template`, body);
  }

  // TODO: type body
  createCustomApplication(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${environment.backendApiUrl}/application/custom`, body);
  }

  // TODO: type body
  createTemplate(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${environment.backendApiUrl}/template`, body);
  }

  deleteTemplate(id: number): Observable<any> {
    return this.http.delete<any>(`${environment.backendApiUrl}/template/${id}`);
  }

  getApplications(): Observable<Application[]> {
    return this.http.get<Application[]>(`${environment.backendApiUrl}/application`);
  }

  getApplicationById(id: number): Observable<TemplateApplication | CustomApplication> {
    return this.http.get<TemplateApplication | CustomApplication>(
      `${environment.backendApiUrl}/application/${id}`
    );
  }

  approveApplicationById(id: number): Observable<null> {
    return this.http.patch<any>(`${environment.backendApiUrl}/application/approve/${id}`, null);
  }

  denyApplicationById(id: number): Observable<null> {
    return this.http.patch<any>(`${environment.backendApiUrl}/application/deny/${id}`, null);
  }

  getAllNamespaces(): Observable<Namespace[]> {
    return this.http.get<Namespace[]>(`${environment.backendApiUrl}/namespace?all=1`);
  }

  saveNamespace(namespace: NamespaceSave): Observable<void> {
    return this.http.put<void>(`${environment.backendApiUrl}/namespace`, namespace);
  }

  getAllUser(): Observable<User[]> {
    return this.http.get<User[]>(`${environment.backendApiUrl}/user`);
  }
}
