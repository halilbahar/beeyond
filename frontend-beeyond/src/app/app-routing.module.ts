import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './core/guards/auth.guard';

const routes: Routes = [
  {
    path: '', canActivate: [AuthGuard], children: [
      {path: 'dashboard', loadChildren: () => import('./modules/dashboard/dashboard.module').then(m => m.DashboardModule)},
      {path: 'blueprint', loadChildren: () => import('./modules/blueprint/blueprint.module').then(m => m.BlueprintModule)},
      {path: 'profile', loadChildren: () => import('./modules/profile/profile.module').then(m => m.ProfileModule)},
      {path: 'accounting', loadChildren: () => import('./modules/accounting/accounting.module').then(m => m.AccountingModule)},
      {path: 'management', loadChildren: () => import('./modules/management/management.module').then(m => m.ManagementModule)},
      {path: 'template', loadChildren: () => import('./modules/template/template.module').then(m => m.TemplateModule)}
    ]
  },
  {path: 'login', loadChildren: () => import('./modules/login/login.module').then(m => m.LoginModule)}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {preloadingStrategy: PreloadAllModules})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
