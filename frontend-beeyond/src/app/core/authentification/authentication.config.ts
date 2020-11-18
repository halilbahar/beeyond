import { AuthConfig } from 'angular-oauth2-oidc';

export const authConfig: AuthConfig = {
  issuer: 'http://localhost:8180/auth/realms/oauth2-realm',

  redirectUri: window.location.origin,

  clientId: 'oauth-client',

  responseType: 'code',

  scope: 'offline_access',

  showDebugInformation: true,
};
