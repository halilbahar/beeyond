import { Component, HostBinding, HostListener, Input, OnInit } from '@angular/core';
import { AUTO_STYLE, animate, state, style, transition, trigger } from '@angular/animations';
import { ActivatedRoute, Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { ConstraintEditDialogComponent } from '../constraint-edit-dialog/constraint-edit-dialog.component';
import { Constraint } from 'src/app/shared/models/constraint.model';

const DEFAULT_DURATION = 300;

@Component({
  selector: 'app-constraint-detail',
  templateUrl: './constraint-detail.component.html',
  styleUrls: ['./constraint-detail.component.scss'],
  animations: [
    trigger('collapse', [
      state('false', style({ height: AUTO_STYLE, visibility: AUTO_STYLE })),
      state('true', style({ height: '0', visibility: 'hidden' })),
      transition('false => true', animate(DEFAULT_DURATION + 'ms ease-in')),
      transition('true => false', animate(DEFAULT_DURATION + 'ms ease-out'))
    ])
  ]
})
export class ConstraintDetailComponent implements OnInit {
  @Input() title: string;
  @Input() description: string;
  @Input() type: string;
  @Input() isKubernetesObject: boolean;
  @Input() constraint?: Constraint;

  @HostBinding('class.constraint') hasConstraint = false;
  @HostBinding('class.disabled') isDisabled = false;

  collapsed = true;

  constructor(private router: Router, private route: ActivatedRoute, private dialog: MatDialog) {}

  ngOnInit(): void {
    const {enum: enumArray, regex, min, max, disabled} = this.constraint || {};
    this.hasConstraint = this.constraint != null && (enumArray || regex || min || max) != null;
    this.isDisabled = disabled;
  }

  openEditDialog(): void {
    const path = this.route.snapshot.url.map(segment => segment.path).join('/') + '/' + this.title;

    const dialogRef = this.dialog.open(ConstraintEditDialogComponent, {
      autoFocus: false,
      minWidth: '50%',
      data: { type: this.type, path, constraint: this.constraint }
    });

    dialogRef.afterClosed().subscribe(cancelled => {
      if (!cancelled) {
        this.navigate('./');
      }
    });
  }

  navigate(path: string): void {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate([path], { relativeTo: this.route });
  }
}
