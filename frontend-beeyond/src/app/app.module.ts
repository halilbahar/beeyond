import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { CoreModule } from './core/core.module';
import { MonacoEditorModule } from 'ngx-monaco-editor';
import { MatProgressBarModule } from '@angular/material/progress-bar';

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
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
