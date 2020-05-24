import { Component } from '@angular/core';
import { AuthenticationService } from './core/authentification/authentication.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'beeyond';
  isLoggedIn = false;

  constructor(private authenticationService: AuthenticationService) {
    this.authenticationService.getCurrentUser().subscribe(user => this.isLoggedIn = user !== null);
  }
}
