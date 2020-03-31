import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-side-navigation',
  templateUrl: './side-navigation.component.html',
  styleUrls: ['./side-navigation.component.css']
})
export class SideNavigationComponent implements OnInit {

  private agenda = [
    { name: 'Dashboard', icon: 'speed' },
    { name: 'Blueprint', icon: 'list_alt' },
    { name: 'Profile', icon: 'account_circle' },
    { name: 'Accounting', icon: 'account_balance' },
    { name: 'Management', icon: 'desktop_windows' }
  ];

  constructor() { }

  ngOnInit(): void {
  }

}
