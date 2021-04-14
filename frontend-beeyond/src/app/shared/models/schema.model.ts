import { Constraint } from './constraint.model';

export interface Schema {
  description: string;
  required: string[] | null;
  type: string;
  properties: { [key: string]: Property };
  'x-kubernetes-group-version-kind': GroupKindVersion[];
  'x-constraint'?: Constraint;
}

export interface Property {
  description: string;
  type: string;
  format: string | null;
  items: PropertyItem;
  'x-is-kubernetes-object': boolean;
  'x-constraint'?: Constraint;
}

export interface PropertyItem {
  type: string;
}

export interface GroupKindVersion {
  group: string;
  kind: string;
  version: string;
}
