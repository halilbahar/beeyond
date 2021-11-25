import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProfileRoutingModule } from './profile-routing.module';
import { ProfileComponent } from './pages/profile/profile.component';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from '@angular/material/tabs';
import { NamespaceListComponent } from './components/namespace-list/namespace-list.component';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatTreeModule } from '@angular/material/tree';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { ManagementModule } from '../management/management.module';

@NgModule({
  declarations: [ProfileComponent, NamespaceListComponent],
  imports: [
    CommonModule,
    ProfileRoutingModule,
    MatCardModule,
    MatTabsModule,
    MatExpansionModule,
    MatTreeModule,
    MatIconModule,
    MatButtonModule,
    ManagementModule
  ]
})
export class ProfileModule {}
