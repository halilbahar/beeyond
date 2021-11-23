import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProgressBarService {
  hidden = new BehaviorSubject<boolean>(false);

  start(): void {
    this.hidden.next(true);
  }

  finish(): void {
    this.hidden.next(false);
  }
}
