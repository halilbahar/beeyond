export interface Schema {
  description: string;
  required: string[] | null;
  type: string;
  properties: Property;
  'x-kubernetes-group-version-kind': GroupKindVersion;
}

export interface Property {
  description: string;
  type: string;
  format: string | null;
  items: PropertyItem;
  enum: string[] | null;
}

export interface PropertyItem {
  type: string;
}

export interface GroupKindVersion {
  group: string;
  kind: string;
  version: string;
}
