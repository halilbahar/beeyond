import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { Namespace } from '../../../../shared/models/namespace.model';

@Component({
  selector: 'app-blueprint-new',
  templateUrl: './blueprint-new.component.html',
  styleUrls: ['./blueprint-new.component.scss']
})
export class BlueprintNewComponent implements OnInit {
  secondFormGroup: FormGroup;
  thirdFormGroup: FormGroup;
  blueprintType: string;

  monacoOptions = { language: 'yaml', scrollBeyondLastLine: false };

  wildcards: string[] = [];
  templates: Template[] = [];
  template: Template;
  message = '';
  namespaces: Namespace[];
  templateId: number;

  constructor(
    private router: Router,
    private fb: FormBuilder,
    private backendApiService: BackendApiService,
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

    this.thirdFormGroup = this.fb.group({
      name: this.fb.control('', Validators.required),
      description: this.fb.control(''),
      class: this.fb.control('', Validators.required),
      fromDate: this.fb.control(null, Validators.required),
      toDate: this.fb.control(null, Validators.required),
      namespace: this.fb.control('', Validators.required),
      purpose: this.fb.control('', [Validators.required, Validators.maxLength(255)])
    });

    this.refreshNamespaces();
  }

  createBlueprint() {
    const blueprint = {
      ...this.secondFormGroup.value,
      ...this.thirdFormGroup.value
    };
    // this.backendApiService.createCustomApplication()
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

  get fields(): FormArray {
    return this.secondFormGroup.controls.fields as FormArray;
  }

  updateColor(val: string) {
    this.blueprintType = val;
  }

  private refreshNamespaces(): void {
    this.backendApiService.getAllNamespaces().subscribe(namespaces => {
      this.namespaces = namespaces;
    });
  }
}
