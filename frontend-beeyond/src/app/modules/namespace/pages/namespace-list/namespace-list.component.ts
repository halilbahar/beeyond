import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { Namespace } from 'src/app/shared/models/namespace.model';
import { NamespaceDialogComponent } from '../../components/namespace-dialog/namespace-dialog.component';

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

  createNamespaceDialog(): void {
    const dialogRef = this.dialog.open(NamespaceDialogComponent, { width: '50vw', autoFocus: false });
    
    dialogRef.afterClosed().subscribe(saved => {
      if (saved) {
        this.refreshNamespaces();
      }
    });
  }

  editNamespaceDialog(namespace: Namespace): void {
    this.dialog.open(NamespaceDialogComponent, {
      width: '50vw',
      autoFocus: false,
      data: namespace
    });
  }

  private refreshNamespaces(): void {
    this.backendApiService.getNamespaces().subscribe(namespaces => {
      this.namespaces = namespaces;
    });
  }
}
