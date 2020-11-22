import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Template } from '../shared/models/template.model';

@Injectable({
  providedIn: 'root'
})
export class HttpApiService {
  constructor(private httpClient: HttpClient) {}

  public getAllTemplates(): Observable<Template[]> {
    return this.httpClient.get<Template[]>(environment.apiUrl + '/template');
  }

  public createCustomApplication(payload) {
    return this.httpClient.post(environment.apiUrl + '/custom-application', payload);
  }
}
