import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlueprintComponent } from './blueprint.component';
import { BlueprintTemplateComponent } from './blueprint-template/blueprint-template.component';
import { BlueprintRoutingModule } from './blueprint-routing.module';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@NgModule({
  declarations: [
    BlueprintComponent,
    BlueprintTemplateComponent
  ],
  imports: [
    CommonModule,
    BlueprintRoutingModule,
    MatCardModule,
    MatDividerModule,
    MatListModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule
  ]
})
export class BlueprintModule { }
