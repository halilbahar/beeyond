import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Namespace } from 'src/app/shared/models/namespace.model';

@Component({
  selector: 'app-namespace-delete-dialog',
  templateUrl: './namespace-delete-dialog.component.html',
  styleUrls: ['./namespace-delete-dialog.component.scss']
})
export class NamespaceDeleteDialogComponent {
  constructor(@Inject(MAT_DIALOG_DATA) public namespace: Namespace) {}
}
