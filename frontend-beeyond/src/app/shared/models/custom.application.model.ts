import { User } from './user.model';

export interface CustomApplication {
  content: string;
  note: string;
  status: string;
  user: User;
}
