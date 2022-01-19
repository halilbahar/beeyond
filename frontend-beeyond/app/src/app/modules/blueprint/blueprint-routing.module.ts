import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BlueprintComponent } from './pages/blueprint/blueprint.component';
import { BlueprintTemplateComponent } from './pages/blueprint-template/blueprint-template.component';
import { BlueprintNewComponent } from './pages/blueprint-new/blueprint-new.component';

const routes: Routes = [
  { path: '', component: BlueprintComponent },
  { path: 'template/:id', component: BlueprintTemplateComponent },
  { path: 'new', component: BlueprintNewComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BlueprintRoutingModule {}
