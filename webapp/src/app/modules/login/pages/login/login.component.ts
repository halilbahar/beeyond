import { Component, OnInit } from '@angular/core';
import { AuthenticationService } from '../../../../core/authentification/authentication.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  public loginForm: FormGroup;
  private returnUrl: string;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private authenticationService: AuthenticationService,
    private fb: FormBuilder
  ) {
    if (this.authenticationService.currentUserValue()) {
      this.router.navigate(['/']).then(console.log);
    }
  }

  ngOnInit(): void {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });

    this.returnUrl = this.route.snapshot.queryParams.returnUrl || '/';
  }

  submitUserLogin() {
    if (!this.loginForm.valid) {
      return;
    }

    const controls = this.loginForm.controls;
    this.authenticationService.login(controls.username.value, controls.password.value)
      .pipe(first())
      .subscribe(data => this.router.navigate([this.returnUrl]));
  }
}
