import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-side-navigation',
  templateUrl: './side-navigation.component.html',
  styleUrls: ['./side-navigation.component.scss']
})
export class SideNavigationComponent implements OnInit {
  @Input() name = 'default';

  agenda = [
    { name: 'Dashboard', icon: 'speed', router: '/dashboard' },
    { name: 'Blueprint', icon: 'list_alt', router: '/blueprint' },
    { name: 'Profile', icon: 'account_circle', router: '/profile' },
    { name: 'Accounting', icon: 'account_balance', router: '/accounting' },
    { name: 'Management', icon: 'desktop_windows', router: '/management' },
    { name: 'Template', icon: 'bakery_dining', router: '/template' }
  ];

  constructor() {}

  ngOnInit(): void {}
}
