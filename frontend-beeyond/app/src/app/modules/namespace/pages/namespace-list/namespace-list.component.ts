import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { NamespaceSave } from 'src/app/shared/models/namespace-save.model';
import { Namespace } from 'src/app/shared/models/namespace.model';
import { NamespaceDeleteDialogComponent } from '../../components/namespace-delete-dialog/namespace-delete-dialog.component';
import { NamespaceDialogComponent } from '../../components/namespace-dialog/namespace-dialog.component';
import { AuthenticationService } from '../../../../core/authentification/authentication.service';

@Component({
  selector: 'app-namespace-list',
  templateUrl: './namespace-list.component.html',
  styleUrls: ['./namespace-list.component.scss']
})
export class NamespaceListComponent implements OnInit {
  namespaces: Namespace[];

  constructor(private backendApiService: BackendApiService, private dialog: MatDialog) {}

  ngOnInit(): void {
    this.refreshNamespaces();
  }

  openCreateNamespaceDialog(): void {
    const dialogRef = this.dialog.open(NamespaceDialogComponent, {
      width: '50vw',
      autoFocus: false
    });

    dialogRef.afterClosed().subscribe(saved => {
      if (saved) {
        this.refreshNamespaces();
      }
    });
  }

  openDeleteNamespaceDialog(namespace: Namespace): void {
    const deletedNamespace: NamespaceSave = {
      namespace: namespace.namespace,
      users: []
    };

    const dialogRef = this.dialog.open(NamespaceDeleteDialogComponent, {
      width: '50vw',
      autoFocus: false,
      data: namespace
    });

    dialogRef.afterClosed().subscribe(shouldDelete => {
      if (shouldDelete) {
        this.backendApiService
          .saveNamespace(deletedNamespace)
          .subscribe(() => this.refreshNamespaces());
      }
    });
  }

  openEditNamespaceDialog(namespace: Namespace): void {
    this.dialog.open(NamespaceDialogComponent, {
      width: '50vw',
      autoFocus: false,
      data: namespace
    });
  }

  private refreshNamespaces(): void {
    this.backendApiService.getAllNamespaces().subscribe(namespaces => {
      this.namespaces = namespaces;
    });
  }
}
