import { ENTER, COMMA } from '@angular/cdk/keycodes';
import { Component, Inject, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';
import { MatChipInputEvent } from '@angular/material/chips';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Observable } from 'rxjs';
import { startWith, map } from 'rxjs/operators';
import { BackendApiService } from 'src/app/core/services/backend-api.service';
import { NamespaceSave } from 'src/app/shared/models/namespace-save.model';
import { Namespace } from 'src/app/shared/models/namespace.model';

@Component({
  selector: 'app-namespace-dialog',
  templateUrl: './namespace-dialog.component.html',
  styleUrls: ['./namespace-dialog.component.scss']
})
export class NamespaceDialogComponent implements OnInit {
  separatorKeysCodes: number[] = [ENTER, COMMA];
  nameControl = new FormControl(null, [
    Validators.required,
    Validators.maxLength(253),
    Validators.pattern(/^[a-z0-9]([-a-z0-9]*[a-z0-9])?$/)
  ]);
  userControl = new FormControl();
  users: string[] = [];
  filteredUsers: Observable<string[]>;
  allUsers: string[] = [];

  constructor(
    private backendApiService: BackendApiService,
    private dialogRef: MatDialogRef<NamespaceDialogComponent>,
    @Inject(MAT_DIALOG_DATA) private namespaceToEdit: Namespace | null
  ) {
    if (namespaceToEdit !== null) {
      this.nameControl.setValue(namespaceToEdit.namespace);
      this.nameControl.disable();
      this.users = namespaceToEdit.users.map(user => user.name);
    }
  }

  private get remainingUser(): string[] {
    return this.allUsers.filter(user => !this.users.includes(user));
  }

  ngOnInit(): void {
    this.backendApiService.getAllUser().subscribe(users => {
      this.allUsers = users.map(user => user.name);
      this.filteredUsers = this.userControl.valueChanges.pipe(
        startWith(null),
        map((user: string | null) => (user ? this.filter(user) : this.remainingUser.slice()))
      );
    });
  }

  add(event: MatChipInputEvent): void {
    const input = event.input;
    const value = event.value;

    const trimmedValue = (value || '').trim();
    if (
      trimmedValue &&
      !this.users.includes(trimmedValue) &&
      this.allUsers.includes(trimmedValue)
    ) {
      this.users.push(trimmedValue);
    }

    if (input) {
      input.value = '';
    }
    this.userControl.setValue(null);
  }

  remove(user: string): void {
    const index = this.users.indexOf(user);

    if (index >= 0) {
      this.users.splice(index, 1);
    }
  }

  selected(event: MatAutocompleteSelectedEvent): void {
    const value = event.option.viewValue;
    if (!this.users.includes(value)) {
      this.users.push(value);
    }

    this.userControl.setValue(null);
  }

  save(): void {
    const namespace: NamespaceSave = {
      namespace: this.nameControl.value,
      users: this.users
    };

    this.backendApiService.saveNamespace(namespace).subscribe(() => this.dialogRef.close(true));
  }

  private filter(value: string): string[] {
    const filterValue = value.toLowerCase();
    return this.remainingUser.filter(user => user.indexOf(filterValue) === 0);
  }
}
