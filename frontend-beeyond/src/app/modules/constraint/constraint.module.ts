import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConstraintRoutingModule } from './constraint-routing.module';
import { RootConstraintComponent } from './pages/root-constraint/root-constraint.component';
import { MatCardModule } from '@angular/material/card';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [RootConstraintComponent],
  imports: [CommonModule, ConstraintRoutingModule, MatCardModule, MatExpansionModule, MatIconModule]
})
export class ConstraintModule {}
