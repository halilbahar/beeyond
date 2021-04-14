import { User } from './user.model';

export interface Namespace {
  id?: number;
  namespace: string;
  users?: User[];
  // Frontend
  label: string;
}
