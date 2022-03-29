import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ActivationEnd, NavigationEnd, Params, Router } from '@angular/router';
import { Breadcrumb } from 'src/app/shared/models/breadcrumb.model';
import { AuthenticationService } from '../authentification/authentication.service';
import { SidenavService } from '../services/sidenav.service';
import { ThemeService } from '../services/theme.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  breadcrumbs: Breadcrumb[] = [];

  constructor(
    private router: Router,
    private sidenavService: SidenavService,
    public oAuthService: AuthenticationService
  ) {}

  ngOnInit(): void {
    let segements: { path: string; params: Params }[] = [];

    this.router.events.subscribe(event => {
      if (event instanceof ActivationEnd) {
        const route = event.snapshot;
        segements.unshift({
          path: route.routeConfig.path,
          params: route.params
        });
      }

      if (event instanceof NavigationEnd) {
        const breadcrumbs: Breadcrumb[] = [];
        let link = '';
        for (const segment of segements) {
          let path = segment.path;
          if (path === '') {
            continue;
          }

          if (path === '**') {
            const remainingUrl = event.url.replace(link, '').slice(1);
            const remainingUrlArray = remainingUrl.split('/');

            if (remainingUrl === '') {
              continue;
            }

            for (const remainingSegment of remainingUrlArray) {
              link += `/${remainingSegment}`;
              breadcrumbs.push({ link, title: remainingSegment });
            }
          } else {
            // Remove any path param associated with the path: /path/:id => /path
            const title = path.replace(/\/:\w+/, '');

            // Itterate over the object that contains the path param and
            // replace the placeholder with the actual value: /path/:id => /path/1
            const params = segment.params;
            for (const key in params) {
              if (params.hasOwnProperty(key)) {
                const value = segment.params[key];
                path = path.replace(`:${key}`, value);
              }
            }

            link += `/${path}`;
            breadcrumbs.push({ link, title });
          }
        }

        breadcrumbs[breadcrumbs.length - 1].link = '';
        this.breadcrumbs = breadcrumbs;
        segements = [];
      }
    });
  }

  navigateFromBreadcrumb(url: string): void {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate([url]);
  }

  toggleSideNavigation(): void {
    const minimized = this.sidenavService.minimized;
    minimized.next(!minimized.value);
  }

  logout(): void {
    this.oAuthService.logOut();
  }
}
