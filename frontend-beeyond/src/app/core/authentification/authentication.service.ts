import { Injectable } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { authConfig } from './authentication.config';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  username = new BehaviorSubject<string>('');

  constructor(private oAuthService: OAuthService) {}

  login() {
    this.oAuthService.initLoginFlow();
  }

  initializeLogin(): void {
    this.oAuthService.configure(authConfig);
    this.oAuthService.setupAutomaticSilentRefresh();
    this.oAuthService.loadDiscoveryDocumentAndTryLogin().then(_ => {
      if (!this.oAuthService.hasValidAccessToken()) {
        this.oAuthService.initLoginFlow();
      } else {
        this.oAuthService
          .loadUserProfile()
          .then(profile => this.username.next(profile.preferred_username));
      }
    });
  }
}
