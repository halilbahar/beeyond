import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SideNavigationComponent } from './side-navigation/side-navigation.component';
import { HeaderComponent } from './header/header.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { OAuthModule } from 'angular-oauth2-oidc';
import { environment } from '../../environments/environment';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  declarations: [SideNavigationComponent, HeaderComponent],
  exports: [SideNavigationComponent, HeaderComponent],
  imports: [
    CommonModule,
    RouterModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    OAuthModule.forRoot({
      resourceServer: {
        allowedUrls: [environment.backendApiUrl],
        sendAccessToken: true
      }
    }),
    MatButtonModule
  ]
})
export class CoreModule {}
