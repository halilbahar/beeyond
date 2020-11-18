import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SideNavigationComponent } from './side-navigation/side-navigation.component';
import { HeaderComponent } from './header/header.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { OAuthModule } from 'angular-oauth2-oidc';

@NgModule({
  declarations: [
    SideNavigationComponent,
    HeaderComponent
  ],
  exports: [
    SideNavigationComponent,
    HeaderComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    OAuthModule.forRoot()
  ]
})
export class CoreModule {}
