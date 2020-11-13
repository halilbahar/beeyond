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

  constructor(private router: Router) { }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.breadcrumbs = []
        const links = event.url.split('/');
        links.shift();

        links.forEach(link => {
          this.breadcrumbs.push({
            link: link,
            title: link.charAt(0).toUpperCase() + link.substr(1)
          })
        });

        this.headerTitle = this.breadcrumbs[0].title;
        console.log(this.breadcrumbs);
      }
    });
  }

  toggleSideNavigation(): void {
    // TODO: close / open sidenav
  }
}
