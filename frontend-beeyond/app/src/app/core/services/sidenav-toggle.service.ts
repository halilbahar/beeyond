import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SidenavToggleService {
  minimized = new BehaviorSubject<boolean>(false);
}
