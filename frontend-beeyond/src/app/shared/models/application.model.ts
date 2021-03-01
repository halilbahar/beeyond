import { ApplicationStatus } from './application-status.enum';
import { User } from './user.model';

export abstract class Application {
  id: number;
  status: ApplicationStatus;
  note: string;
  owner: User;
  createdAt: string;
}