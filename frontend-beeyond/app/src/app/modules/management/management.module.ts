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
import { ApplicationPreviewDialogComponent } from './components/application-preview-dialog/application-preview-dialog.component';
import { MatDialogModule } from '@angular/material/dialog';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { ApplicationContentComponent } from './components/application-content/application-content.component';
import { MatTabsModule } from '@angular/material/tabs';
import { ApplicationAttributesComponent } from './components/application-attributes/application-attributes.component';

@NgModule({
  declarations: [
    ApplicationComponent,
    ApplicationReviewComponent,
    VariableListComponent,
    ApplicationPreviewDialogComponent,
    ApplicationContentComponent,
    ApplicationAttributesComponent
  ],
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
    MatSortModule,
    MatDialogModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatSnackBarModule,
    MatTabsModule
  ],
  exports: [ApplicationContentComponent, ApplicationReviewComponent]
})
export class ManagementModule {}
