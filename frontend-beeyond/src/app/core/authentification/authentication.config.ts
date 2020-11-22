import { AuthConfig } from 'angular-oauth2-oidc';

export const authConfig: AuthConfig = {
  issuer: 'http://localhost:8180/auth/realms/school',
  redirectUri: window.location.origin,
  clientId: 'beeyond-spa',
  responseType: 'code',
  scope: 'offline_access',
  showDebugInformation: true
};
