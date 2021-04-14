import { Component, Inject, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  FormControl,
  FormGroup,
  ValidationErrors,
  ValidatorFn
} from '@angular/forms';
import { MatChipInputEvent } from '@angular/material/chips';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';
import { Constraint } from 'src/app/shared/models/constraint.model';
import { COMMA, ENTER } from '@angular/cdk/keycodes';

@Component({
  selector: 'app-constraint-edit-dialog',
  templateUrl: './constraint-edit-dialog.component.html',
  styleUrls: ['./constraint-edit-dialog.component.scss']
})
export class ConstraintEditDialogComponent implements OnInit {
  form: FormGroup;
  readonly separatorKeysCodes: number[] = [ENTER, COMMA];
  private disablingForm = false;

  constructor(
    @Inject(MAT_DIALOG_DATA)
    public data: { path: string; type: string; constraint: Constraint | null },
    private dialogRef: MatDialogRef<ConstraintEditDialogComponent>,
    private fb: FormBuilder,
    private validationApiService: ValidationApiService
  ) {}

  ngOnInit(): void {
    const { enum: enumArray, min, max, regex } = this.data.constraint || {};
    const group: { [key: string]: any } = {};
    if (this.isRegexAllowed()) {
      group.regex = [regex || null];
    }
    if (this.isEnumAllowed()) {
      group.enum = [enumArray || []];
    }
    if (this.isMinMaxAllowed()) {
      group.min = [min || null];
      group.max = [max || null];
    }

    this.form = this.fb.group(group);
    this.form.setValidators(this.validConstraint());
    // If the controls have inti
    const valueChange = this.onValueChange();
    valueChange();

    // Disable other controls when 1 control has a value
    this.form.valueChanges.subscribe(valueChange);
  }

  get enumControl(): FormControl {
    return this.form.get('enum') as FormControl;
  }

  createConstraint(): void {
    this.validationApiService
      .createConstraint(this.data.path, this.form.value as Constraint)
      .subscribe(() => this.dialogRef.close());
  }

  deleteConstraint(): void {
    this.validationApiService
      .deleteConstraint(this.data.path)
      .subscribe(() => this.dialogRef.close());
  }

  addEnum(event: MatChipInputEvent): void {
    const input = event.input;
    const value = event.value;

    if (value?.trim()) {
      const newValue = this.enumControl.value;
      newValue.push(value.trim());
      this.enumControl.setValue(newValue);
    }

    if (input) {
      input.value = '';
    }
  }

  removeEnum(index: number): void {
    const newValue = [...this.enumControl.value];
    newValue.splice(index, 1);
    this.enumControl.setValue(newValue);
  }

  private isRegexAllowed(): boolean {
    return this.data.type === 'string';
  }

  private isEnumAllowed(): boolean {
    return this.data.type === 'string' || this.data.type === 'integer';
  }

  private isMinMaxAllowed(): boolean {
    return this.data.type === 'integer';
  }

  private onValueChange(): () => void {
    return () => {
      if (!this.disablingForm) {
        this.disablingForm = true;

        if ((this.form.controls.regex?.value || '') !== '') {
          this.form.controls.enum?.disable();
        } else if (this.form.controls.enum?.value.length > 0) {
          this.form.controls.regex?.disable();
          this.form.controls.min?.disable();
          this.form.controls.max?.disable();
        } else if (
          (this.form.controls.min?.value || '') !== '' ||
          (this.form.controls.max?.value || '') !== ''
        ) {
          this.form.controls.enum?.disable();
        } else {
          this.form.enable();
        }
        this.disablingForm = false;
      }
    };
  }

  private validConstraint(): ValidatorFn {
    return (control: AbstractControl): ValidationErrors | null => {
      const { regex = '', enum: enumArray = '', min, max } = control.value;

      const valid = regex !== '' || enumArray?.length > 0 || min != null || max != null;
      if (valid) {
        return;
      }

      return { constraintValid: valid };
    };
  }
}
