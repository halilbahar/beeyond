import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { map } from 'rxjs/operators';
import { User } from '../../shared/models/user.model';
import { OAuthService } from 'angular-oauth2-oidc';
import { authConfig } from './authentication.config';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private currentUserSubject: BehaviorSubject<User>;
  private readonly currentUser: Observable<User>;

  constructor(private oAuthService: OAuthService) {
    this.currentUserSubject = new BehaviorSubject(JSON.parse(localStorage.getItem('currentUser')));
    this.currentUser = this.currentUserSubject.asObservable();
  }

  /*login(username: string, password: string) {
    return this.http.post<User>(`${environment.apiUrl}/authentication/login`, {username, password})
      .pipe(map(user => {
        localStorage.setItem('currentUser', JSON.stringify(user));
        this.currentUserSubject.next(user);
        return user;
      }));
  }*/

  login() {
    this.oAuthService.initLoginFlow();
  }

  logout() {
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
  }

  initializeLogin(): void {
    this.oAuthService.configure(authConfig);
    this.oAuthService.setupAutomaticSilentRefresh();
    this.oAuthService.loadDiscoveryDocumentAndTryLogin().then(_ => {
      if (!this.oAuthService.hasValidAccessToken()) {
        this.oAuthService.initLoginFlow();
      }
    });
  }

  getCurrentUser() {
    return this.currentUser;
  }

  currentUserValue() {
    return this.currentUserSubject.value;
  }
}
