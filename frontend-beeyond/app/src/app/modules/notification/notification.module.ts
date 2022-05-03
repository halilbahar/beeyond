import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NotificationRoutingModule } from './notification-routing.module';
import { NotificationComponent } from './pages/notification/notification.component';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';

@NgModule({
  declarations: [NotificationComponent],
  imports: [CommonModule, NotificationRoutingModule, MatIconModule, MatCardModule]
})
export class NotificationModule {}
