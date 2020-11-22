import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  headerTitle = '';
  breadcrumbs = [];

  constructor(private router: Router) {}

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.breadcrumbs = [];
        const url = event.url.split('?')[0];
        const links = url.split('/');
        let curLink = '';
        links.shift();

        links.forEach(link => {
          curLink += '/' + link;
          this.breadcrumbs.push({
            link: curLink,
            title: link.charAt(0).toUpperCase() + link.substr(1)
          });
        });

        this.breadcrumbs[this.breadcrumbs.length - 1].link = '';
        this.headerTitle = this.breadcrumbs[0].title;
      }
    });
  }

  toggleSideNavigation(): void {
    // TODO: close / open sidenav
  }
}
