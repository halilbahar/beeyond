export interface Constraint {
  min: number | null;
  max: number | null;
  enum: string[] | null;
  regex: string | null;
}
