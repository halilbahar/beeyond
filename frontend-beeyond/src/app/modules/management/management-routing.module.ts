import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ApplicationComponent } from './pages/application/application.component';
import { ApplicationReviewComponent } from './pages/application-review/application-review.component';
import { ApplicationReviewResolver } from 'src/app/core/resolver/application-review.resolver';

const routes: Routes = [
  { path: '', component: ApplicationComponent },
  {
    path: 'review/:id',
    component: ApplicationReviewComponent,
    resolve: { application: ApplicationReviewResolver }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ApplicationRoutingModule {}
