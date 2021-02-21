import { Component, Input } from '@angular/core';
import { AUTO_STYLE, animate, state, style, transition, trigger } from '@angular/animations';
import { ActivatedRoute, Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { ConstraintEditDialogComponent } from '../constraint-edit-dialog/constraint-edit-dialog.component';

const DEFAULT_DURATION = 300;

@Component({
  selector: 'app-constraint-detail',
  templateUrl: './constraint-detail.component.html',
  styleUrls: ['./constraint-detail.component.scss'],
  animations: [
    trigger('collapse', [
      state('false', style({ height: AUTO_STYLE, visibility: AUTO_STYLE, padding: AUTO_STYLE })),
      state('true', style({ height: '0', visibility: 'hidden', padding: '0' })),
      transition('false => true', animate(DEFAULT_DURATION + 'ms ease-in')),
      transition('true => false', animate(DEFAULT_DURATION + 'ms ease-out'))
    ])
  ]
})
export class ConstraintDetailComponent {
  @Input() title: string;
  @Input() description: string;
  @Input() type: string;
  @Input() hasRef: boolean;

  collapsed = true;

  constructor(private router: Router, private route: ActivatedRoute, private dialog: MatDialog) {}

  openEditDialog(): void {
    const path = this.route.snapshot.url
      .map(segment => segment.path)
      .reduce((previous, current) => previous + current, '') + '/' + this.title;

    this.dialog.open(ConstraintEditDialogComponent, {
      autoFocus: false,
      minWidth: '50%',
      data: { type: this.type, path }
    });
  }

  navigate(path: string): void {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate([path], { relativeTo: this.route });
  }
}
