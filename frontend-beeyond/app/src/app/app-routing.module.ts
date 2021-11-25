import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { PageNotFoundComponent } from './shared/components/page-not-found/page-not-found.component';
import { AuthGuard } from './core/guards/auth.guard';

const routes: Routes = [
  {
    path: '',
    canActivateChild: [AuthGuard],
    children: [
      {
        path: '',
        redirectTo: '/blueprint',
        pathMatch: 'full'
      },
      {
        path: 'accounting',
        loadChildren: () =>
          import('./modules/accounting/accounting.module').then(m => m.AccountingModule)
      },
      {
        path: 'management',
        loadChildren: () =>
          import('./modules/management/management.module').then(m => m.ManagementModule)
      },
      {
        path: 'template',
        loadChildren: () => import('./modules/template/template.module').then(m => m.TemplateModule)
      },
      {
        path: 'namespace',
        loadChildren: () =>
          import('./modules/namespace/namespace.module').then(m => m.NamespaceModule)
      },
      {
        path: 'constraint',
        loadChildren: () =>
          import('./modules/constraint/constraint.module').then(m => m.ConstraintModule)
      }
    ]
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
