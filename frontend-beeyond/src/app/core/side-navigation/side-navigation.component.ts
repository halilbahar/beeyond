import { Component, OnInit } from '@angular/core';
import { AUTO_STYLE, state, style, trigger } from '@angular/animations';
import { SidenavToggleService } from '../services/sidenav-toggle.service';
import { AuthenticationService } from '../authentification/authentication.service';

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
  agenda = [
    { name: 'Dashboard', icon: 'speed', router: '/dashboard' },
    { name: 'Blueprint', icon: 'list_alt', router: '/blueprint' },
    { name: 'Profile', icon: 'account_circle', router: '/profile' },
    { name: 'Accounting', icon: 'account_balance', router: '/accounting' },
    { name: 'Management', icon: 'desktop_windows', router: '/management' },
    { name: 'Template', icon: 'bakery_dining', router: '/template' }
  ];

  constructor(
    public sidenavToggleService: SidenavToggleService,
    public authenticationService: AuthenticationService
  ) {}

  ngOnInit(): void {}
}
