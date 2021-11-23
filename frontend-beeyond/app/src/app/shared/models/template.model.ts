import { TemplateField } from './template-field.model';

export interface Template {
  id: number;
  name: string;
  description: string;
  content: string;
  fields: TemplateField[];
  deleted: boolean;
  namespace: string;
}
