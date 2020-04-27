import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpApiService {

  constructor(private httpClient: HttpClient) { }

  public getAllTemplates(): Observable<string[]> {
    return this.httpClient.get<string[]>(environment.apiUrl + '/template');
  }
}

