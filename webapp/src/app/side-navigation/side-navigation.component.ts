import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-side-navigation',
  templateUrl: './side-navigation.component.html',
  styleUrls: ['./side-navigation.component.css']
})
export class SideNavigationComponent implements OnInit {

  agenda = [
    { name: 'Dashboard', icon: 'speed' },
    { name: 'Blueprint', icon: 'list_alt' },
    { name: 'Profile', icon: 'account_circle' },
    { name: 'Accounting', icon: 'account_balance' },
    { name: 'Management', icon: 'desktop_windows' }
  ];

  @Input()
  name = 'default';

  constructor() { }

  ngOnInit(): void {
  }

}
