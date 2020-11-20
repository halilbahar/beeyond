import { Component } from '@angular/core';
import { AuthenticationService } from './core/authentification/authentication.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'beeyond';

  constructor(private authenticationService: AuthenticationService) {
    this.authenticationService.initializeLogin();
  }
}
