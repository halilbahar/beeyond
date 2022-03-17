import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SideNavigationComponent } from './side-navigation/side-navigation.component';
import { HeaderComponent } from './header/header.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { OAuthModule } from 'angular-oauth2-oidc';
import { MatButtonModule } from '@angular/material/button';
import { authModuleConfig } from './authentification/oauth-module.config';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatMenuModule } from '@angular/material/menu';
import { FormsModule } from '@angular/forms';
import { BaseComponent } from './services/base.component';

@NgModule({
  declarations: [SideNavigationComponent, HeaderComponent, BaseComponent],
  exports: [SideNavigationComponent, HeaderComponent, BaseComponent],
  imports: [
    CommonModule,
    RouterModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    OAuthModule.forRoot(authModuleConfig),
    MatButtonModule,
    MatSlideToggleModule,
    MatMenuModule,
    FormsModule
  ]
})
export class CoreModule {}
