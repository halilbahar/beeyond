import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApplicationComponent } from './pages/application/application.component';
import { ApplicationRoutingModule } from './management-routing.module';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { ApplicationReviewComponent } from './pages/application-review/application-review.component';
import { VariableListComponent } from './components/variable-list/variable-list.component';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatSelectModule } from '@angular/material/select';
import { MatSortModule } from '@angular/material/sort';
import { TemplateApplicationPreviewDialogComponent } from './components/template-application-preview-dialog/template-application-preview-dialog.component';

@NgModule({
  declarations: [ApplicationComponent, ApplicationReviewComponent, VariableListComponent, TemplateApplicationPreviewDialogComponent],
  imports: [
    CommonModule,
    ApplicationRoutingModule,
    MatCardModule,
    MatTableModule,
    MonacoEditorModule,
    MatButtonModule,
    FormsModule,
    MatIconModule,
    MatInputModule,
    MatAutocompleteModule,
    ReactiveFormsModule,
    MatSelectModule,
    MatSortModule
  ]
})
export class ManagementModule {}
