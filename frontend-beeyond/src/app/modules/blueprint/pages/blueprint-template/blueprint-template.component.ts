import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Template } from '../../../../shared/models/template.model';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-blueprint-template',
  templateUrl: './blueprint-template.component.html',
  styleUrls: ['./blueprint-template.component.scss']
})
export class BlueprintTemplateComponent implements OnInit {
  template: Template;
  templateForm: FormGroup;
  id: number;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private backendApiService: BackendApiService,
    private fb: FormBuilder,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    this.id = Number(this.route.snapshot.paramMap.get('id'));
    this.backendApiService.getTemplateById(this.id).subscribe(template => {
      this.template = template;

      const fieldValues = [];
      this.template.fields.forEach(field => fieldValues.push(this.createFieldValue(field.id)));

      this.templateForm = this.fb.group({
        templateId: [this.id],
        note: ['', Validators.maxLength(255)],
        fieldValues: this.fb.array(fieldValues)
      });
    });
  }

  createFieldValue(fieldId: number): FormGroup {
    return this.fb.group({
      value: ['', Validators.required],
      fieldId: [fieldId]
    });
  }

  submitApplication(): void {
    this.backendApiService.createTemplateApplication(this.templateForm.value).subscribe(
      () => {
        this.router.navigate(['dashboard']).then(navigated => {
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
        this.snackBar.open(error.error.map(err => err.message).join('\n'), 'close', {
          duration: undefined
        });
      }
    );
  }
}
