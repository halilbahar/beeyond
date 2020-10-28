import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TemplateComponent } from './pages/template/template.component';
import { TemplateCreateComponent } from './pages/template-create/template-create.component';

const routes: Routes = [
  {path: '', component: TemplateComponent},
  {path: 'create', component: TemplateCreateComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TemplateRoutingModule { }
