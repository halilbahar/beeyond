import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { AUTO_STYLE, state, style, trigger } from '@angular/animations';
import { SidenavService } from '../services/sidenav.service';
import { AuthenticationService } from '../authentification/authentication.service';
import { config } from '../config/user-role.config';
import { BackendApiService } from '../services/backend-api.service';
import { ThemeService } from '../services/theme.service';
import { Notification } from '../../shared/models/notification.model';
import { NotificationStatus } from '../../shared/models/notification-status.enum';
import { MediaMatcher } from '@angular/cdk/layout';
import { BaseComponent } from '../services/base.component';

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
export class SideNavigationComponent extends BaseComponent implements OnInit {
  isDarkTheme: boolean;
  agenda = [
    { name: 'Blueprint', icon: 'list_alt', router: '/blueprint' },
    { name: 'Profile', icon: 'account_circle', router: '/profile' },
    /*{
      name: 'Accounting',
      icon: 'account_balance',
      router: '/accounting',
      requiredRoles: [config.adminRole]
    },*/
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
  lastAccess = new Date(Number(localStorage.getItem('lastAccess')));

  constructor(
    public sidenavService: SidenavService,
    public authenticationService: AuthenticationService,
    public backendApiService: BackendApiService,
    private themeService: ThemeService,
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher
  ) {
    super(changeDetectorRef, media);
    this.isDarkTheme = themeService.isDarkTheme.value;
    this.changes.subscribe(() => {
      this.sidenavService.minimized.next(!this.mobileQuery?.matches);
    });
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
        this.backendApiService.getNotifications().subscribe(notifications => {
          // Only show last 3 notifications in the side nav -> the rest of the notifications is available in a separate page
          this.notifications = notifications
            .sort((n1, n2) => new Date(n1.createdAt).getTime() - new Date(n2.createdAt).getTime())
            .reverse()
            .slice(0, 3);
          this.notifications.forEach(n => (n.createdAt = new Date(n.createdAt)));
        });
      }
    });
  }

  toggleTheme() {
    window.localStorage.setItem('isDarkTheme', String(this.isDarkTheme));
    this.themeService.isDarkTheme.next(this.isDarkTheme);
  }

  logOut(): void {
    this.authenticationService.logOut();
  }
}
