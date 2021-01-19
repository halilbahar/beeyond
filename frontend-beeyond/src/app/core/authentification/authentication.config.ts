import { AuthConfig } from 'angular-oauth2-oidc';
import { environment } from 'src/environments/environment';

export const authConfig: AuthConfig = {
  issuer: environment.identityProviderIssuer,
  redirectUri: window.location.origin,
  clientId: 'beeyond-spa',
  responseType: 'code',
  scope: 'offline_access',
  showDebugInformation: true
};
