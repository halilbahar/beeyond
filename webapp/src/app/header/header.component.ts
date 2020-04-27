import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  headerTitle = '';

  constructor(private router: Router) { }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        const title = event.url.substr(1);
        this.headerTitle = title.charAt(0).toUpperCase() + title.substr(1);
      }
    });
  }

  toggleSideNavigation(): void {
    // TODO: close / open sidenav
  }
}
