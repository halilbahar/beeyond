import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RootConstraintComponent } from './pages/root-constraint/root-constraint.component';

const routes: Routes = [{ path: '', component: RootConstraintComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ConstraintRoutingModule {}
