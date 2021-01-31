import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConstraintRoutingModule } from './constraint-routing.module';
import { RootConstraintComponent } from './pages/root-constraint/root-constraint.component';

@NgModule({
  declarations: [RootConstraintComponent],
  imports: [CommonModule, ConstraintRoutingModule]
})
export class ConstraintModule {}
