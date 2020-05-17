import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Template } from '../modules/blueprint/template.model';

@Injectable({
  providedIn: 'root'
})
export class HttpApiService {

  constructor(private httpClient: HttpClient) { }

  public getAllTemplates(): Observable<Template[]> {
    return this.httpClient.get<Template[]>(environment.apiUrl + '/template');
  }
}

