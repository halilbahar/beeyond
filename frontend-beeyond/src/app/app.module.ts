import { BrowserModule } from '@angular/platform-browser';
import { APP_INITIALIZER, NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { CoreModule } from './core/core.module';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { ConfigService } from './core/services/config.service';
import { AuthenticationService } from './core/authentification/authentication.service';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    HttpClientModule,
    CoreModule,
    MonacoEditorModule.forRoot(),
    MatProgressBarModule
  ],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: (auth: AuthenticationService, config: ConfigService) => () =>
        config.init().then(() => auth.initializeLogin().then()),
      deps: [AuthenticationService, ConfigService],
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
