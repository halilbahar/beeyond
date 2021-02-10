import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-template-create',
  templateUrl: './template-create.component.html',
  styleUrls: ['./template-create.component.scss']
})
export class TemplateCreateComponent implements OnInit {
  firstFormGroup: FormGroup;
  secondFormGroup: FormGroup;
  thirdFormGroup: FormGroup;

  monacoOptions = { language: 'yaml', scrollBeyondLastLine: false };

  wildcards: string[] = [];

  constructor(
    private router: Router,
    private fb: FormBuilder,
    private backendApiService: BackendApiService,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    this.firstFormGroup = this.fb.group({
      content: ['', Validators.required]
    });

    this.secondFormGroup = this.fb.group({
      fields: this.fb.array([])
    });

    this.thirdFormGroup = this.fb.group({
      name: ['', [Validators.required, Validators.maxLength(255)]],
      description: ['', Validators.maxLength(255)]
    });

    this.firstFormGroup.controls.content.valueChanges.subscribe(content => {
      this.wildcards = [];
      this.fields.clear();

      const regex = /%([\w]+)%/g;
      let match;

      do {
        match = regex.exec(content);
        if (match) {
          this.wildcards.push(match[1]);
          this.fields.push(this.createWildcardField(match[1]));
        }
      } while (match);
    });
  }

  createTemplate(): void {
    const template = {
      ...this.firstFormGroup.value,
      ...this.secondFormGroup.value,
      ...this.thirdFormGroup.value
    };
    this.backendApiService.createTemplate(template).subscribe(() => {
      this.router.navigate(['template']).then(navigated => {
        if (navigated) {
          this.snackBar.open('Your Template was created', 'close', {
            duration: undefined
          });
        }
      });
    });
  }

  createWildcardField(wildcard: string): FormGroup {
    return this.fb.group({
      wildcard: [wildcard],
      label: ['', [Validators.required, Validators.maxLength(255)]],
      description: ['', Validators.maxLength(255)]
    });
  }

  get fields(): FormArray {
    return this.secondFormGroup.controls.fields as FormArray;
  }
}
