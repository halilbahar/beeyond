import { TemplateFieldValue } from './template-field-value.model';
import { TemplateField } from './template-field.model';
import { User } from './user.model';

export interface TemplateApplication {
    note: string;
    fieldValues: TemplateFieldValue[] //TODO;
    id: number;
    status: string;
    templateId: number;
    //owner: User; funktioniert nicht mit
}
