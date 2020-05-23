import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlueprintComponent } from './pages/blueprint/blueprint.component';
import { BlueprintTemplateComponent } from './pages/blueprint-template/blueprint-template.component';
import { BlueprintRoutingModule } from './blueprint-routing.module';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatTabsModule } from '@angular/material/tabs';
import { FormsModule } from '@angular/forms';

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
    MatInputModule,
    MatTabsModule,
    FormsModule
  ]
})
export class BlueprintModule { }
