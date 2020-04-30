import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { BlueprintComponent } from './blueprint/blueprint.component';
import { AccountingComponent } from './accounting/accounting.component';
import { ProfileComponent } from './profile/profile.component';
import { ManagementComponent } from './management/management.component';
import { BlueprintTemplateComponent } from './blueprint/blueprint-template/blueprint-template.component';


const routes: Routes = [
  {path: 'dashboard', component: DashboardComponent},
  {path: 'blueprint', component: BlueprintComponent},
  {path: 'blueprint/:name', component: BlueprintTemplateComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'accounting', component: AccountingComponent},
  {path: 'management', component: ManagementComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
