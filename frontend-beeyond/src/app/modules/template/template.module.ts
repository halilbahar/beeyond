import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TemplateRoutingModule } from './template-routing.module';
import { TemplateComponent } from './pages/template/template.component';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { SharedModule } from '../../shared/shared.module';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { TemplateCreateComponent } from './pages/template-create/template-create.component';
import { ReactiveFormsModule } from '@angular/forms';
import { MatDividerModule } from '@angular/material/divider';
import { MatStepperModule } from '@angular/material/stepper';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { TemplateDetailComponent } from './pages/template-detail/template-detail.component';

@NgModule({
  declarations: [TemplateComponent, TemplateCreateComponent, TemplateDetailComponent],
  imports: [
    CommonModule,
    TemplateRoutingModule,
    MatCardModule,
    MatButtonModule,
    SharedModule,
    MatInputModule,
    MatIconModule,
    ReactiveFormsModule,
    MatDividerModule,
    MatStepperModule,
    MatSnackBarModule,
    MonacoEditorModule
  ]
})
export class TemplateModule {}
