import { Component, OnInit } from '@angular/core';
import { Template } from '../../../../shared/models/template.model';
import { ApiService } from '../../../../core/services/api.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.scss']
})
export class BlueprintComponent implements OnInit {
  templates: Template[] = [];
  customApplicationForm: FormGroup;
  monacoOptions = { language: 'yaml', scrollBeyondLastLine: false };
  message = '';

  constructor(
    private router: Router,
    private apiService: ApiService,
    private fb: FormBuilder,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    this.apiService.getTemplates().subscribe(templates => {
      this.templates = templates.filter(template => !template.deleted);
      if (this.templates.length === 0) {
        this.message =
          'There are currently no templates available. Wait until your administrator creates one.';
      }
    });

    this.customApplicationForm = this.fb.group({
      content: ['', Validators.required],
      note: ['', Validators.maxLength(255)]
    });
  }

  sendCustomTemplate(): void {
    this.apiService.createCustomApplication(this.customApplicationForm.value).subscribe(() => {
      this.router.navigate(['dashboard']).then(navigated => {
        if (navigated) {
          this.snackBar.open(
            'Your application was sent will be reviewed as soon as possible',
            'close',
            { duration: undefined }
          );
        }
      });
    });
  }
}
