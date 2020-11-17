import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApplicationComponent } from './pages/application/application.component';
import { ApplicationRoutingModule } from './application-routing.module';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { ApplicationReviewComponent } from './pages/application-review/application-review.component';
import { VariableListComponent } from './components/variable-list/variable-list.component';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  declarations: [
    ApplicationComponent,
    ApplicationReviewComponent,
    VariableListComponent
  ],
  imports: [
    CommonModule,
    ApplicationRoutingModule,
    MatCardModule,
    MatTableModule,
    MonacoEditorModule,
    MatButtonModule,
    FormsModule
  ]
})
export class ApplicationModule { }
