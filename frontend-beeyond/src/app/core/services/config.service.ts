import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs/operators';
import { Config } from '../../shared/models/config.model';
import { authModuleConfig } from '../authentification/oauth-module.config';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {
  config: Config;

  constructor(private httpClient: HttpClient) {}

  init(): Promise<any> {
    return this.httpClient
      .get<Config>('/assets/config.json')
      .pipe(
        tap(config => {
          this.config = config;
          authModuleConfig.resourceServer.allowedUrls.push(config.backendApiUrl);
        })
      )
      .toPromise();
  }
}
