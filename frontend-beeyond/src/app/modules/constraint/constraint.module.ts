import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConstraintRoutingModule } from './constraint-routing.module';
import { ConstraintComponent } from './pages/constraint/constraint.component';
import { MatCardModule } from '@angular/material/card';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [ConstraintComponent],
  imports: [CommonModule, ConstraintRoutingModule, MatCardModule, MatExpansionModule, MatIconModule]
})
export class ConstraintModule {}
