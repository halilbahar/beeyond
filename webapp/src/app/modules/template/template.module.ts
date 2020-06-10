import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TemplateRoutingModule } from './template-routing.module';
import { TemplateComponent } from './pages/template/template.component';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { SharedModule } from '../../shared/shared.module';

@NgModule({
  declarations: [TemplateComponent],
  imports: [
    CommonModule,
    TemplateRoutingModule,
    MatCardModule,
    MatButtonModule,
    SharedModule,
  ]
})
export class TemplateModule { }
