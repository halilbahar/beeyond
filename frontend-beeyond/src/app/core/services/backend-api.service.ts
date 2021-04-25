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

@Injectable({
  providedIn: 'root'
})
export class BackendApiService {
  constructor(private http: HttpClient, private configService: ConfigService) {}

  getTemplates(): Observable<Template[]> {
    return this.http.get<Template[]>(`${this.configService.config.backendApiUrl}/template`);
  }

  getTemplateById(id: number): Observable<Template> {
    return this.http.get<Template>(`${this.configService.config.backendApiUrl}/template/${id}`);
  }

  // TODO: type body
  createTemplateApplication(body: Record<string, unknown>): Observable<null> {
    return this.http.post<null>(
      `${this.configService.config.backendApiUrl}/application/template`,
      body
    );
  }

  // TODO: type body
  createCustomApplication(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(
      `${this.configService.config.backendApiUrl}/application/custom`,
      body
    );
  }

  // TODO: type body
  createTemplate(body: Record<string, unknown>): Observable<any> {
    return this.http.post<any>(`${this.configService.config.backendApiUrl}/template`, body);
  }

  deleteTemplate(id: number): Observable<void> {
    return this.http.delete<void>(`${this.configService.config.backendApiUrl}/template/${id}`);
  }

  getApplications(): Observable<Application[]> {
    return this.http.get<Application[]>(`${this.configService.config.backendApiUrl}/application`);
  }

  getApplicationById(id: number): Observable<TemplateApplication | CustomApplication> {
    return this.http.get<TemplateApplication | CustomApplication>(
      `${this.configService.config.backendApiUrl}/application/${id}`
    );
  }

  approveApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(
      `${this.configService.config.backendApiUrl}/application/approve/${id}`,
      null
    );
  }

  denyApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(
      `${this.configService.config.backendApiUrl}/application/deny/${id}`,
      null
    );
  }

  stopApplicationById(id: number): Observable<void> {
    return this.http.patch<void>(
      `${this.configService.config.backendApiUrl}/application/stop/${id}`,
      null
    );
  }

  getAllNamespaces(): Observable<Namespace[]> {
    return this.http.get<Namespace[]>(`${this.configService.config.backendApiUrl}/namespace?all=1`);
  }

  saveNamespace(namespace: NamespaceSave): Observable<void> {
    return this.http.put<void>(`${this.configService.config.backendApiUrl}/namespace`, namespace);
  }

  getAllUser(): Observable<User[]> {
    return this.http.get<User[]>(`${this.configService.config.backendApiUrl}/user`);
  }
}
