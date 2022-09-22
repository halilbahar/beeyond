import { Component, OnInit } from '@angular/core';
import { BackendApiService } from '../../../../core/services/backend-api.service';
import { Notification } from '../../../../shared/models/notification.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.scss']
})
export class NotificationComponent implements OnInit {
  notifications: Notification[];

  constructor(private backendApiService: BackendApiService, private router: Router) {}

  ngOnInit(): void {
    this.refreshNotifications();
  }

  deleteNotification(notification: Notification) {
    this.backendApiService.deleteNotifications(notification.id).subscribe(() => {
      this.refreshNotifications();
    });
  }

  private refreshNotifications(): void {
    this.backendApiService.getNotifications().subscribe(notifications => {
      this.notifications = notifications.sort(
        (n1, n2) => new Date(n1.createdAt).getTime() - new Date(n2.createdAt).getTime()
      );
      console.log(notifications);
    });
  }
}
