import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AccountingComponent } from './pages/accounting/accounting.component';
import { AccountingRoutingModule } from './accounting-routing.module';

@NgModule({
  declarations: [AccountingComponent],
  imports: [CommonModule, AccountingRoutingModule]
})
export class AccountingModule {}
