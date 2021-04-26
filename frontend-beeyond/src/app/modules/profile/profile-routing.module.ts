import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ApplicationResolver } from 'src/app/core/resolver/application.resolver';
import { ProfileComponent } from './pages/profile/profile.component';

const routes: Routes = [{ path: '', component: ProfileComponent, resolve: { applications: ApplicationResolver} }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProfileRoutingModule {}
