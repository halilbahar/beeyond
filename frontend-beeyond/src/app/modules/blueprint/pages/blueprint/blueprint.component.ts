import { Component, OnInit } from '@angular/core';
import { Template } from '../../../../shared/models/template.model';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Namespace } from '../../../../shared/models/namespace.model';
import { AuthenticationService } from 'src/app/core/authentification/authentication.service';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.scss']
})
export class BlueprintComponent implements OnInit {
  templates: Template[] = [];
  namespaces: Namespace[] = [];
  customApplicationForm: FormGroup;
  monacoOptions = { language: 'yaml', scrollBeyondLastLine: false };
  message = '';

  constructor(
    private router: Router,
    private backendApiService: BackendApiService,
    private fb: FormBuilder,
    private snackBar: MatSnackBar,
    private authenticationService: AuthenticationService
  ) {}

  ngOnInit(): void {
    this.backendApiService.getTemplates().subscribe(templates => {
      this.templates = templates.filter(template => !template.deleted);
      if (this.templates.length === 0) {
        this.message =
          'There are currently no templates available. Wait until your administrator creates one.';
      }
    });

    this.backendApiService.getAllNamespaces().subscribe(namespaces => {
      const defaultNamespace = {
        namespace: this.authenticationService.username.value,
        label: 'Default'
      };

      this.namespaces = namespaces.map(namespace => ({ ...namespace, label: namespace.namespace }));
      this.namespaces.push(defaultNamespace);
      this.customApplicationForm.patchValue({
        namespace: defaultNamespace.namespace
      });
    });

    this.customApplicationForm = this.fb.group({
      content: ['', Validators.required],
      note: ['', Validators.maxLength(255)],
      namespace: ['', Validators.required]
    });
  }

  sendCustomTemplate(): void {
    this.backendApiService.createCustomApplication(this.customApplicationForm.value).subscribe(
      () => {
        this.router.navigate(['dashboard']).then(navigated => {
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
        this.snackBar.open(error.error.map(err => err.message).join('\n'), 'close', {
          duration: undefined
        });
      }
    );
  }
}
