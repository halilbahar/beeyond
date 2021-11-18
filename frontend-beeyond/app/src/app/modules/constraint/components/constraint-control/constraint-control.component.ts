import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-constraint-control',
  templateUrl: './constraint-control.component.html',
  styleUrls: ['./constraint-control.component.scss']
})
export class ConstraintControlComponent {
  @Output() controlChange: EventEmitter<ConstraintControlChange> = new EventEmitter();

  searchValue = '';
  hideDeleted = false;

  sendChanges(): void {
    this.controlChange.emit({ search: this.searchValue, hideDeleted: this.hideDeleted });
  }
}

export type ConstraintControlChange = { search: string; hideDeleted: boolean };
