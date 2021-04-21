import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NameSpaceRoutingModule } from './namespace-routing.module';
import { MatCardModule } from '@angular/material/card';
import { NamespaceListComponent } from './pages/namespace-list/namespace-list.component';
import { MatDialogModule } from '@angular/material/dialog';
import { NamespaceDialogComponent } from './components/namespace-dialog/namespace-dialog.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatChipsModule } from '@angular/material/chips';
import { ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { NamespaceDeleteDialogComponent } from './components/namespace-delete-dialog/namespace-delete-dialog.component';

@NgModule({
  declarations: [NamespaceListComponent, NamespaceDialogComponent, NamespaceDeleteDialogComponent],
  imports: [
    CommonModule,
    NameSpaceRoutingModule,
    MatCardModule,
    MatDialogModule,
    MatFormFieldModule,
    MatChipsModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatButtonModule
  ]
})
export class NamespaceModule {}
