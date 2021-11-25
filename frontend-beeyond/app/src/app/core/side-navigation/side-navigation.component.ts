import { Component, OnInit } from '@angular/core';
import { AUTO_STYLE, state, style, trigger } from '@angular/animations';
import { SidenavToggleService } from '../services/sidenav-toggle.service';
import { AuthenticationService } from '../authentification/authentication.service';
import { config } from '../config/user-role.config';
import { BackendApiService } from '../services/backend-api.service';
import { ThemeService } from '../services/theme.service';
import { Notification } from '../../shared/models/notification.model';
import { NotificationStatus } from '../../shared/models/notification-status.enum';

@Component({
  selector: 'app-side-navigation',
  templateUrl: './side-navigation.component.html',
  styleUrls: ['./side-navigation.component.scss'],
  animations: [
    trigger('minimized', [
      state('false', style({ width: AUTO_STYLE, display: AUTO_STYLE })),
      state('true', style({ width: '0', display: 'none' }))
    ])
  ]
})
export class SideNavigationComponent implements OnInit {
  theme: boolean;
  agenda = [
    { name: 'Blueprint', icon: 'list_alt', router: '/blueprint' },
    { name: 'Profile', icon: 'account_circle', router: '/profile' },
    {
      name: 'Accounting',
      icon: 'account_balance',
      router: '/accounting',
      requiredRoles: [config.adminRole]
    },
    {
      name: 'Management',
      icon: 'desktop_windows',
      router: '/management',
      requiredRoles: [config.adminRole]
    },
    {
      name: 'Template',
      icon: 'bakery_dining',
      router: '/template',
      requiredRoles: [config.adminRole]
    },
    {
      name: 'Namespace',
      icon: 'dns',
      router: '/namespace',
      requiredRoles: [config.adminRole]
    },
    {
      name: 'Constraint',
      icon: 'attach_file',
      router: '/constraint',
      requiredRoles: [config.adminRole]
    }
  ];
  actualAgenda = [];
  notifications: Notification[] = [];
  notificationStatus = NotificationStatus;

  constructor(
    public sidenavToggleService: SidenavToggleService,
    public authenticationService: AuthenticationService,
    public backendApiService: BackendApiService,
    private themeService: ThemeService
  ) {
    this.theme = themeService.theme.value;
  }

  ngOnInit(): void {
    this.authenticationService.roles.subscribe(async res => {
      if (this.agenda) {
        this.actualAgenda = this.agenda.filter(item => {
          let found = true;
          if (item.requiredRoles) {
            item.requiredRoles.forEach(
              requiredRole => (found = !!res.find(role => role === requiredRole))
            );
          }
          return found;
        });
        this.backendApiService
          .getNotifications()
          .subscribe(notifications => (this.notifications = notifications));
      }
    });
  }

  toggleTheme() {
    this.themeService.theme.next(this.theme);
  }

  logOut(): void {
    this.authenticationService.logOut();
  }
}
