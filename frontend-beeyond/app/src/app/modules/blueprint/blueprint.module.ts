import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlueprintComponent } from './pages/blueprint/blueprint.component';
import { BlueprintTemplateComponent } from './pages/blueprint-template/blueprint-template.component';
import { BlueprintRoutingModule } from './blueprint-routing.module';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatTabsModule } from '@angular/material/tabs';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { MatSelectModule } from '@angular/material/select';

@NgModule({
  declarations: [BlueprintComponent, BlueprintTemplateComponent],
  imports: [
    CommonModule,
    BlueprintRoutingModule,
    MatCardModule,
    MatDividerModule,
    MatListModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatTabsModule,
    FormsModule,
    SharedModule,
    ReactiveFormsModule,
    MatSnackBarModule,
    MonacoEditorModule,
    MatSelectModule
  ]
})
export class BlueprintModule {}
