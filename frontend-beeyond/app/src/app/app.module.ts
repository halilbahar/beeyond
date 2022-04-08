import { BrowserModule } from '@angular/platform-browser';
import { APP_INITIALIZER, NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { CoreModule } from './core/core.module';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { ConfigService } from './core/services/config.service';
import { AuthenticationService } from './core/authentification/authentication.service';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { HttpRequestInterceptor } from './core/interceptors/http-request.interceptor';

@NgModule({
  declarations: [AppComponent],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        AppRoutingModule,
        HttpClientModule,
        CoreModule,
        MonacoEditorModule.forRoot(),
        MatProgressBarModule,
        MatCardModule,
        MatIconModule,
        MatProgressSpinnerModule
    ],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: (auth: AuthenticationService, config: ConfigService) => () =>
        config.init().then(() => auth.initializeLogin()),
      deps: [AuthenticationService, ConfigService],
      multi: true
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpRequestInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
