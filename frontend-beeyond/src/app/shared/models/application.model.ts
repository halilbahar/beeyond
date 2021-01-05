import { ApplicationStatus } from './application-status.enum';
import { User } from './user.model';

export abstract class Application {
  status: ApplicationStatus;
  note: string;
  owner: User;
}
