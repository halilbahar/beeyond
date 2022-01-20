import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { Namespace } from '../../../../shared/models/namespace.model';
import { DatePipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthenticationService } from '../../../../core/authentification/authentication.service';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.scss']
})
export class BlueprintComponent implements OnInit {
  secondFormGroup: FormGroup;
  thirdFormGroup: FormGroup;
  blueprintType = '';

  monacoOptions = { language: 'yaml', scrollBeyondLastLine: false };

  wildcards: string[] = [];
  templates: Template[] = [];
  template: Template;
  message = '';
  namespaces: Namespace[];
  templateId: number = null;
  templateForm: FormGroup;

  constructor(
    public authenticationService: AuthenticationService,
    private router: Router,
    private fb: FormBuilder,
    private snackBar: MatSnackBar,
    private backendApiService: BackendApiService
  ) {}

  ngOnInit(): void {
    this.backendApiService.getTemplates().subscribe(templates => {
      this.templates = templates.filter(template => !template.deleted);
      if (this.templates.length === 0) {
        this.message =
          'There are currently no templates available. Wait until your administrator creates one.';
      }
    });

    this.secondFormGroup = this.fb.group({
      content: ['', Validators.required]
    });

    this.secondFormGroup.controls.content.valueChanges.subscribe(content => {
      this.wildcards = [];
      if (this.fields) {
        this.fields.clear();
      }

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

    this.thirdFormGroup = this.fb.group({
      note: this.fb.control(''),
      class: this.fb.control('', Validators.required),
      to: this.fb.control(null, Validators.required),
      namespace: this.fb.control('', Validators.required),
      purpose: this.fb.control('', [Validators.required, Validators.maxLength(255)])
    });

    this.refreshNamespaces();
  }

  stepperSelectionChange(event) {
    switch (event.selectedIndex) {
      case 1:
        if (this.blueprintType === 'Template') {
          this.loadTemplate();
        }
        break;
    }
  }

  createBlueprint() {
    if (this.blueprintType === 'Custom') {
      const blueprint = {
        ...this.secondFormGroup.value,
        ...this.thirdFormGroup.value
      };

      blueprint.to = new DatePipe('en-US').transform(blueprint.to, 'dd.MM.yyyy');

      this.backendApiService.createCustomApplication(blueprint).subscribe(
        () => {
          this.router.navigate(['/profile']).then(navigated => {
            if (navigated) {
              this.snackBar.open(
                'Your application was sent will be reviewed as soon as possible',
                'close',
                { duration: undefined }
              );
            }
          });
        },
        error => {
          this.snackBar.open(
            error.error.map(err => err.message + ' - ' + err.key).join('\n'),
            'close',
            {
              duration: undefined,
              panelClass: ['new-line']
            }
          );
        }
      );
    } else if (this.blueprintType === 'Template') {
      const blueprint = {
        ...this.templateForm.value,
        ...this.thirdFormGroup.value
      };

      blueprint.to = new DatePipe('en-US').transform(blueprint.to, 'dd.MM.yyyy');

      this.backendApiService.createTemplateApplication(blueprint).subscribe(
        () => {
          this.router.navigate(['/profile']).then(navigated => {
            if (navigated) {
              this.snackBar.open(
                'Your application was sent will be reviewed as soon as possible',
                'close',
                { duration: 2000 }
              );
            }
          });
        },
        error => {
          this.snackBar.open(
            error.error.map(err => err.message + ' - ' + err.key).join('\n'),
            'close',
            {
              duration: undefined,
              panelClass: ['new-line']
            }
          );
        }
      );
    }
  }

  createWildcardField(wildcard: string): FormGroup {
    return this.fb.group({
      wildcard: [wildcard],
      label: ['', [Validators.required, Validators.maxLength(255)]],
      description: ['', Validators.maxLength(255)]
    });
  }

  setTemplateId(id: number) {
    this.templateId = id;
  }

  loadTemplate() {
    if (this.blueprintType === 'Template' && this.templateId) {
      if (this.template && this.templateId === this.template.id) {
        return;
      }
      this.backendApiService.getTemplateById(this.templateId).subscribe(template => {
        this.template = template;

        const fieldValues = [];
        this.template.fields.forEach(field => fieldValues.push(this.createFieldValue(field.id)));

        this.templateForm = this.fb.group({
          templateId: [this.templateId],
          note: ['', Validators.maxLength(255)],
          fieldValues: this.fb.array(fieldValues)
        });
      });
    }
  }

  createFieldValue(fieldId: number): FormGroup {
    return this.fb.group({
      value: ['', Validators.required],
      fieldId: [fieldId]
    });
  }

  get fields(): FormArray {
    return this.secondFormGroup.controls.fields as FormArray;
  }

  updateColor(val: string) {
    this.blueprintType = val;
  }

  updateBlueprintType(val: string) {
    if (val === 'Custom') {
      this.templateId = null;
    }
    this.blueprintType = val;
  }

  private refreshNamespaces(): void {
    this.backendApiService.getUserNamespaces().subscribe(namespaces => {
      const defaultNamespace = {
        namespace: this.authenticationService.username.value,
        label: 'Default'
      };

      this.namespaces = namespaces
        .map(namespace => ({ ...namespace, label: namespace.namespace }))
        .filter(namespace => namespace.namespace !== this.authenticationService.username.value);
      this.namespaces.push(defaultNamespace);
    });
  }
}
