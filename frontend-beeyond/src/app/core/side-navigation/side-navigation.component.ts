import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-side-navigation',
  templateUrl: './side-navigation.component.html',
  styleUrls: ['./side-navigation.component.scss']
})
export class SideNavigationComponent implements OnInit {

  agenda = [
    { name: 'Dashboard', icon: 'speed', router: '/dashboard' },
    { name: 'Blueprint', icon: 'list_alt', router: '/blueprint' },
    { name: 'Application', icon: 'book_online', router: '/application' },
    { name: 'Profile', icon: 'account_circle', router: '/profile' },
    { name: 'Accounting', icon: 'account_balance', router: '/accounting' },
    { name: 'Management', icon: 'desktop_windows', router: '/management' },
    { name: 'Template', icon: 'bakery_dining', router: '/template' }
  ];

  @Input()
  name = 'default';

  @Input()
  isLoggedIn = false;

  constructor() { }

  ngOnInit(): void {
  }

}
