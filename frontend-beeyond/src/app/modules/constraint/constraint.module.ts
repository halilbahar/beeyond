import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ConstraintRoutingModule } from './constraint-routing.module';
import { RootConstraintComponent } from './pages/root-constraint/root-constraint.component';
import { MatCardModule } from '@angular/material/card';
import { MatExpansionModule } from '@angular/material/expansion';
import { ConstraintAccordionComponent } from './components/constraint-accordion/constraint-accordion.component';

@NgModule({
  declarations: [RootConstraintComponent, ConstraintAccordionComponent],
  imports: [CommonModule, ConstraintRoutingModule, MatCardModule, MatExpansionModule]
})
export class ConstraintModule {}
