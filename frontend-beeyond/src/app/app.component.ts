import { Component, HostBinding } from '@angular/core';
import { AuthenticationService } from './core/authentification/authentication.service';
import { ProgressBarService } from './core/services/progress-bar.service';
import { animate, AUTO_STYLE, state, style, transition, trigger } from '@angular/animations';
import { ThemeService } from './core/services/theme.service';
import { OverlayContainer } from '@angular/cdk/overlay';

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
  @HostBinding('class') className = '';
  title = 'beeyond';

  constructor(
    public authenticationService: AuthenticationService,
    public progressBarService: ProgressBarService,
    public themeService: ThemeService,
    private overlayContainer: OverlayContainer
  ) {
    this.themeService.theme.subscribe(value => {
      this.className = value ? 'darkMode' : '';
      if (value) {
        this.overlayContainer.getContainerElement().classList.add('darkMode');
      } else {
        this.overlayContainer.getContainerElement().classList.remove('darkMode');
      }
    });
  }
}
