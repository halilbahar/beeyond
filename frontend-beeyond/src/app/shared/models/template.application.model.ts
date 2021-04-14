import { Application } from './application.model';
import { TemplateFieldValue } from './template-field-value.model';

export interface TemplateApplication extends Application {
  note: string;
  fieldValues: TemplateFieldValue[];
  id: number;
  templateId: number;
}
