import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TemplateListComponent } from './components/template-list/template-list.component';
import { MatListModule } from '@angular/material/list';
import { MatRippleModule } from '@angular/material/core';
import { MatIconModule } from '@angular/material/icon';


@NgModule({
  declarations: [
    TemplateListComponent
  ],
  exports: [
    TemplateListComponent
  ],
  imports: [
    CommonModule,
    MatListModule,
    MatRippleModule,
    MatIconModule
  ]
})
export class SharedModule { }
