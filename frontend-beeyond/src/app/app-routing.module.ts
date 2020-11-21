import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './core/guards/auth.guard';
import { ApplicationResolver } from './core/resolver/application.resolver';
import { PageNotFoundComponent } from './shared/components/page-not-found/page-not-found.component';

const routes: Routes = [

  {path: 'dashboard', loadChildren: () => import('./modules/dashboard/dashboard.module').then(m => m.DashboardModule)},
  {path: 'blueprint', loadChildren: () => import('./modules/blueprint/blueprint.module').then(m => m.BlueprintModule)},
  {path: 'profile', loadChildren: () => import('./modules/profile/profile.module').then(m => m.ProfileModule)},
  {path: 'accounting', loadChildren: () => import('./modules/accounting/accounting.module').then(m => m.AccountingModule)},
  {path: 'management', loadChildren: () => import('./modules/management/management.module').then(m => m.ManagementModule)},
  {path: 'template', loadChildren: () => import('./modules/template/template.module').then(m => m.TemplateModule)},
  {
<<<<<<< HEAD
    path: '', canActivate: [AuthGuard], children: [
      { path: 'dashboard', loadChildren: () => import('./modules/dashboard/dashboard.module').then(m => m.DashboardModule) },
      { path: 'blueprint', loadChildren: () => import('./modules/blueprint/blueprint.module').then(m => m.BlueprintModule) },
      { path: 'profile', loadChildren: () => import('./modules/profile/profile.module').then(m => m.ProfileModule) },
      { path: 'accounting', loadChildren: () => import('./modules/accounting/accounting.module').then(m => m.AccountingModule) },
      { path: 'management', loadChildren: () => import('./modules/management/management.module').then(m => m.ManagementModule), resolve: { applications: ApplicationResolver } },
      { path: 'template', loadChildren: () => import('./modules/template/template.module').then(m => m.TemplateModule) }
    ]
=======
    path: 'application',
    loadChildren: () => import('./modules/application/application.module').then(m => m.ApplicationModule),
    resolve: {applications: ApplicationResolver}
>>>>>>> 3c74057c09acdcd471849fc068c543252b4d9cfc
  },
  {path: '', redirectTo: '/dashboard', pathMatch: 'full'},
  {
    path: '**', component: PageNotFoundComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {preloadingStrategy: PreloadAllModules, relativeLinkResolution: 'legacy'})],
  exports: [RouterModule]
})
export class AppRoutingModule {}
