export interface Constraint {
  min?: number;
  max?: number;
  enum?: string[];
  regex?: string;
  disabled?: boolean;
}
