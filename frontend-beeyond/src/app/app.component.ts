import { Component } from '@angular/core';
import { AuthenticationService } from './core/authentification/authentication.service';
import { ProgressBarService } from './core/services/progress-bar.service';
import { animate, AUTO_STYLE, state, style, transition, trigger } from '@angular/animations';
import { combineLatest, forkJoin, from, timer } from 'rxjs';

const DEFAULT_DURATION = 300;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  animations: [
    trigger('fadeOut', [
      state('false', style({ opacity: AUTO_STYLE, visibility: AUTO_STYLE })),
      state('true', style({ opacity: 0, visibility: 'hidden' })),
      transition('false => true', animate(DEFAULT_DURATION + 'ms ease-in')),
      transition('true => false', animate(DEFAULT_DURATION + 'ms ease-out'))
    ])
  ]
})
export class AppComponent {
  title = 'beeyond';
  oidcLoaded = false;

  constructor(
    private authenticationService: AuthenticationService,
    public progressBarService: ProgressBarService
  ) {
    forkJoin([from(this.authenticationService.initializeLogin()), timer(500)]).subscribe(
      () => (this.oidcLoaded = true)
    );
  }
}
