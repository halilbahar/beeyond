import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Router } from '@angular/router';
import { Template } from '../../../../shared/models/template.model';
import { Namespace } from '../../../../shared/models/namespace.model';
import { DatePipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthenticationService } from '../../../../core/authentification/authentication.service';
import { MediaMatcher } from '@angular/cdk/layout';
import { BaseComponent } from '../../../../core/services/base.component';
import { ThemeService } from '../../../../core/services/theme.service';
import * as yaml from 'js-yaml';
import { ApplicationRange } from '../../../../shared/models/application-range.model';
// eslint-disable-next-line max-len
import { ApplicationPreviewDialogComponent } from '../../../management/components/application-preview-dialog/application-preview-dialog.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-blueprint',
  templateUrl: './blueprint.component.html',
  styleUrls: ['./blueprint.component.scss']
})
export class BlueprintComponent extends BaseComponent implements OnInit {
  secondFormGroup: FormGroup;
  thirdFormGroup: FormGroup;
  blueprintType = '';
  wildcards: string[] = [];
  templates: Template[] = [];
  template: Template;
  message = '';
  namespaces: Namespace[];
  templateId: number = null;
  templateForm: FormGroup;
  services = [];

  monacoOptions = {
    language: 'yaml',
    scrollBeyondLastLine: false,
    theme: this.themeService.isDarkTheme.value ? 'vs-dark' : 'vs-light',
    automaticLayout: true
  };

  constructor(
    public authenticationService: AuthenticationService,
    private router: Router,
    private fb: FormBuilder,
    private snackBar: MatSnackBar,
    private backendApiService: BackendApiService,
    private themeService: ThemeService,
    private dialog: MatDialog,
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher
  ) {
    super(changeDetectorRef, media);
    this.themeService.isDarkTheme.subscribe(value => {
      this.monacoOptions = { ...this.monacoOptions, theme: value ? 'vs-dark' : 'vs-light' };
    });
  }

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

    this.secondFormGroup.controls.content.valueChanges.subscribe(value => {
      this.loadServices();
    });

    this.thirdFormGroup = this.fb.group({
      note: this.fb.control('', Validators.required),
      class: this.fb.control('', Validators.required),
      to: this.fb.control(null),
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
    let blueprint = {
      ...this.thirdFormGroup.value
    };

    const temp: any[] = yaml.loadAll(this.getContent());
    temp.map((c: any) => {
      if (c.kind === 'Service' && this.services.find(s => s.name === c.metadata.name).selected) {
        if (!c.metadata.labels) {
          c.metadata.labels = {};
        }
        c.metadata.labels['beeyond-create-ingress'] = 'true';
      }
    });
    blueprint.content = temp.map(c => yaml.dump(c)).join('---\n');

    if (this.blueprintType === 'Custom') {
      blueprint = {
        ...blueprint,
        ...this.secondFormGroup.value,
        content: blueprint.content
      };

      blueprint.to = new DatePipe('en-US').transform(blueprint.to, 'dd.MM.yyyy');

      this.backendApiService.createCustomApplication(blueprint).subscribe(
        () =>
          this.router.navigate(['/profile']).then(navigated => {
            if (navigated) {
              this.snackBar.open(
                'Your application was sent will be reviewed as soon as possible',
                'close',
                {
                  duration: 2000,
                  panelClass: ['mat-drawer-container']
                }
              );
            }
          }),
        error => {
          this.snackBar.open(
            error.error.map(err => err.message + ' - ' + err.key).join('\n'),
            'close',
            {
              duration: undefined,
              panelClass: ['new-line', 'mat-drawer-container']
            }
          );
        }
      );
    } else if (this.blueprintType === 'Template') {
      blueprint = {
        ...blueprint,
        ...this.templateForm.value
      };
      blueprint.to = new DatePipe('en-US').transform(blueprint.to, 'dd.MM.yyyy');

      this.backendApiService.createTemplateApplication(blueprint).subscribe(
        () =>
          this.router.navigate(['/profile']).then(navigated => {
            if (navigated) {
              this.snackBar.open(
                'Your application was sent will be reviewed as soon as possible',
                'close',
                {
                  duration: 2000,
                  panelClass: ['mat-drawer-container']
                }
              );
            }
          }),
        error => {
          this.snackBar.open(
            error?.error?.map(err => err.message + ' - ' + err.key).join('\n'),
            'close',
            {
              duration: undefined,
              panelClass: ['new-line', 'mat-drawer-container']
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
    this.services = [];
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

        this.loadServices();
      });
    }
  }

  getContent() {
    let content = '';
    if (this.blueprintType === 'Template') {
      const regex = /%([\w-]+)%/g;
      let temp = this.template.content;
      this.template.fields.forEach(f => {
        if (this.templateForm.controls.fieldValues.value.find(v => v.fieldId === f.id).value) {
          const re = new RegExp('%' + f.wildcard + '%', 'g');
          temp = temp.replace(
            re,
            this.templateForm.controls.fieldValues.value.find(v => v.fieldId === f.id).value
          );
        }
      });
      content = temp.replace(regex, 'temp');
    } else {
      content = this.secondFormGroup.controls.content.value;
    }
    return content;
  }

  loadServices(): void {
    const content = this.getContent();
    this.services = [];
    try {
      yaml
        .loadAll(content)
        .filter((v: any) => v.kind === 'Service')
        .forEach((v: any) => {
          this.services.push({
            name: v.metadata.name,
            selected: false,
            ports: v.spec.ports.map(p => p.port)
          });
        });
    } catch (e) {}
  }

  createFieldValue(fieldId: number): FormGroup {
    return this.fb.group({
      value: ['', Validators.required],
      fieldId: [fieldId]
    });
  }

  updateColor(val: string) {
    this.blueprintType = val;
  }

  updateBlueprintType(val: string) {
    if (val === 'Custom') {
      this.templateId = null;
    }
    this.blueprintType = val;
    this.services = [];
  }

  openDialog(): void {
    const templateContent = this.template.content;
    const lines = templateContent.split('\n');

    let content = '';
    const ranges: ApplicationRange[] = [];
    const wildcardRegex = /%(.+?)%/g;

    for (let i = 0; i < lines.length; i++) {
      let line = lines[i];
      let match: RegExpExecArray;

      while ((match = wildcardRegex.exec(line)) !== null) {
        const { wildcard, label, description, id } = this.template.fields.find(
          data => data.wildcard === match[0].replace(/%/g, '')
        );
        const value = this.templateForm.controls.fieldValues.value.find(
          f => f.fieldId === id
        ).value;
        if (value) {
          line = line.replace(`%${wildcard}%`, value);
        }

        ranges.push({
          lineNumber: i + 1,
          startColumn: match.index + 1,
          endColumn: match.index + 1 + (value ? value.length : wildcard.length + 2),
          wildcard,
          label,
          description
        });
      }

      content += line + '\n';
    }
    // Remove \n
    content = content.substring(0, content.length - 1);

    this.dialog.open(ApplicationPreviewDialogComponent, {
      data: { content, ranges },
      width: '100%',
      height: '80%',
      autoFocus: false
    });
  }

  private refreshNamespaces(): void {
    this.backendApiService.getUserNamespaces().subscribe(namespaces => {
      this.namespaces = namespaces
        .map(namespace => ({ ...namespace, label: namespace.namespace }))
        .filter(namespace => namespace.namespace !== this.authenticationService.username.value);

      const defaultNamespace = this.namespaces.find(n => n.default);
      defaultNamespace.label = 'Default';
      this.thirdFormGroup.controls.namespace.setValue(defaultNamespace.namespace);
    });
  }
}
