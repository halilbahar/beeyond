import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MediaMatcher } from '@angular/cdk/layout';

@Component({
  selector: 'app-base',
  template: ``,
  styles: []
})
export class BaseComponent {
  mobileQueryListener: () => void;
  mobileQuery: MediaQueryList;

  constructor(
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher
  ) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQuery.onchange = () => changeDetectorRef.detectChanges();
  }
}
