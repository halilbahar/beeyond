import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NamespaceListComponent } from './pages/namespace-list/namespace-list.component';

const routes: Routes = [{ path: '', component: NamespaceListComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class NameSpaceRoutingModule {}
