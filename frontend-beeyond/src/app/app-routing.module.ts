import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { PageNotFoundComponent } from './shared/components/page-not-found/page-not-found.component';

const routes: Routes = [
  {
    path: 'dashboard',
    loadChildren: () => import('./modules/dashboard/dashboard.module').then(m => m.DashboardModule)
  },
  {
    path: 'blueprint',
    loadChildren: () => import('./modules/blueprint/blueprint.module').then(m => m.BlueprintModule)
  },
  {
    path: 'profile',
    loadChildren: () => import('./modules/profile/profile.module').then(m => m.ProfileModule)
  },
  {
    path: 'accounting',
    loadChildren: () =>
      import('./modules/accounting/accounting.module').then(m => m.AccountingModule)
  },
  {
    path: 'template',
    loadChildren: () => import('./modules/template/template.module').then(m => m.TemplateModule)
  },
  {
    path: 'management',
    loadChildren: () =>
      import('./modules/management/management.module').then(m => m.ManagementModule)
  },
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  {
    path: '**',
    component: PageNotFoundComponent
  }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      preloadingStrategy: PreloadAllModules,
      relativeLinkResolution: 'legacy'
    })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule {}
