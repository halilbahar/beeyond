import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { Template } from '../../shared/models/template.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }

  getTemplates() {
    return this.http.get<Template[]>(environment.apiUrl + '/template');
  }
}
