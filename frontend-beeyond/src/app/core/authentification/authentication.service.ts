import { Injectable } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { authConfig } from './authentication.config';
import { BehaviorSubject, identity } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  username = new BehaviorSubject<string>('');
  roles = new BehaviorSubject<string[]>([]);

  constructor(private oAuthService: OAuthService) {}

  logOut() {
    this.oAuthService.logOut();
  }

  async initializeLogin(): Promise<void> {
    this.oAuthService.configure(authConfig);
    await this.oAuthService.loadDiscoveryDocumentAndTryLogin({
      customHashFragment: window.location.search
    });

    if (!this.oAuthService.hasValidAccessToken()) {
      this.oAuthService.initLoginFlow();
    } else {
      this.oAuthService.setupAutomaticSilentRefresh();
      const profile = await this.oAuthService.loadUserProfile();
      this.username.next(profile.preferred_username);
      this.roles.next(this.parseJwt(this.oAuthService.getAccessToken()).realm_access.roles);
    }
  }

  // https://stackoverflow.com/a/38552302/11125147
  private parseJwt(token): any {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    );

    return JSON.parse(jsonPayload);
  }
}
