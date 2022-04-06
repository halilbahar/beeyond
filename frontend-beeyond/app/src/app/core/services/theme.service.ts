import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  isDarkTheme = new BehaviorSubject<boolean>(false);

  constructor() {
    if (window.localStorage.getItem('isDarkTheme') != null) {
      this.isDarkTheme.next(window.localStorage.getItem('isDarkTheme') === 'true');
    } else {
      this.isDarkTheme.next(
        window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
      );
    }
  }
}
