import { User } from './user.model';

export interface CustomApplication {
  content: string;
  id: number;
  note: string;
  status: string;
  //owner: User;
}
