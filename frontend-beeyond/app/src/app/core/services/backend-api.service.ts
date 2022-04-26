import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Template } from '../../shared/models/template.model';
import { Observable } from 'rxjs';
import { Application } from 'src/app/shared/models/application.model';
import { CustomApplication } from 'src/app/shared/models/custom.application.model';
import { TemplateApplication } from 'src/app/shared/models/template.application.model';
import { Namespace } from 'src/app/shared/models/namespace.model';
import { User } from 'src/app/shared/models/user.model';
import { NamespaceSave } from 'src/app/shared/models/namespace-save.model';
import { ConfigService } from './config.service';
import { Notification } from '../../shared/models/notification.model';

@Injectable({
  providedIn: 'root'
})
export class BackendApiService {
  private readonly backendApiUrl: string;

  constructor(private http: HttpClient, private configService: ConfigService) {
    this.backendApiUrl = configService.config.backendApiUrl;
  }

  getTemplates(): Observable<Template[]> {
    return this.http.get<Template[]>(`${this.backendApiUrl}/template`);
  }

  getTemplateById(id: number): Observable<Template> {
    return this.http.get<Template>(`${this.backendApiUrl}/template/${id}`);
  }

  // TODO: type body
  createTemplateApplication(body: Record<string, unknown>): Observable<null> {
    return this.http.post<null>(`${this.backendApiUrl}/application/template`, body);
  }

  // TODO: type body
  createCustomApplication(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${this.backendApiUrl}/application/custom`, body);
  }

  // TODO: type body
  createTemplate(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${this.backendApiUrl}/template`, body);
  }

  deleteTemplate(id: number): Observable<void> {
    return this.http.delete<void>(`${this.backendApiUrl}/template/${id}`);
  }

  getApplications(all: boolean): Observable<Application[]> {
    return this.http.get<Application[]>(`${this.backendApiUrl}/application?all=${all ? 1 : 0}`);
  }

  getApplicationById(id: number): Observable<TemplateApplication | CustomApplication> {
    return this.http.get<TemplateApplication | CustomApplication>(
      `${this.backendApiUrl}/application/${id}`
    );
  }

  approveApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/approve/${id}`, null);
  }

  startApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/start/${id}`, null);
  }

  denyApplicationById(id: number, message: string): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/deny/${id}`, { message });
  }

  stopApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/stop/${id}`, null);
  }

  finishApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/finish/${id}`, null);
  }

  requestApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/request/${id}`, null);
  }

  saveApplicationById(id: number, application: any): Observable<void> {
    return this.http.patch<void>(`${this.backendApiUrl}/application/${id}`, application);
  }

  getAllNamespaces(): Observable<Namespace[]> {
    return this.http.get<Namespace[]>(`${this.backendApiUrl}/namespace?all=1`);
  }

  getUserNamespaces(): Observable<Namespace[]> {
    return this.http.get<Namespace[]>(`${this.backendApiUrl}/namespace`);
  }

  saveNamespace(namespace: NamespaceSave): Observable<void> {
    return this.http.put<void>(`${this.backendApiUrl}/namespace`, namespace);
  }

  getAllUser(): Observable<User[]> {
    return this.http.get<User[]>(`${this.backendApiUrl}/user`);
  }

  getNotifications(): Observable<Notification[]> {
    // Set lastAccess to be able to visibly show that new notifications are available
    window.localStorage.setItem('lastAccess', new Date().getTime().toString());
    return this.http.get<Notification[]>(`${this.backendApiUrl}/notification`);
  }
}
