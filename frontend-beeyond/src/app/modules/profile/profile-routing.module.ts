import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ApplicationReviewResolver } from 'src/app/core/resolver/application-review.resolver';
import { ApplicationResolver } from 'src/app/core/resolver/application.resolver';
import { ApplicationReviewComponent } from '../management/pages/application-review/application-review.component';
import { ProfileComponent } from './pages/profile/profile.component';

const routes: Routes = [
  { path: '', component: ProfileComponent, resolve: { applications: ApplicationResolver } },
  {
    path: 'review/:id',
    component: ApplicationReviewComponent,
    resolve: { application: ApplicationReviewResolver },
    data: { isManagement: false, redirectPath: ['/profile'] }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProfileRoutingModule {}
