import { NotificationStatus } from './notification-status.enum';

export abstract class Notification {
  id: number;
  message: string;
  status: NotificationStatus;
  entityName: string;
  entityId: number;
  userId: number;
  createdAt: Date;
}
