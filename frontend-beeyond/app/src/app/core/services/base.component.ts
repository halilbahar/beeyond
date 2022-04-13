import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MediaMatcher } from '@angular/cdk/layout';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-base',
  template: ``,
  styles: []
})
export abstract class BaseComponent {
  public mobileQuery: MediaQueryList;
  public changes = new BehaviorSubject<boolean>(false);

  protected constructor(changeDetectorRef: ChangeDetectorRef, media: MediaMatcher) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQuery.onchange = () => {
      changeDetectorRef.detectChanges();
      this.changes.next(true);
    };
  }
}
