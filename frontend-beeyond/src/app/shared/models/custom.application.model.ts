import { Application } from './application.model';

export interface CustomApplication extends Application {
  content: string;
  id: number;
  note: string;
}
