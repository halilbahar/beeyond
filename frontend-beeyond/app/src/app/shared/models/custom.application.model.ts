import { Application } from './application.model';

export interface CustomApplication extends Application {
  id: number;
  note: string;
}
