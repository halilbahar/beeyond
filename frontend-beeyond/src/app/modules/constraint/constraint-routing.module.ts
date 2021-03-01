import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ConstraintComponent } from './pages/constraint/constraint.component';

const routes: Routes = [{ path: '**', component: ConstraintComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ConstraintRoutingModule {}
