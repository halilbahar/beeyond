import { Component, EventEmitter, HostBinding, Input, OnInit, Output } from '@angular/core';
import { AUTO_STYLE, animate, state, style, transition, trigger } from '@angular/animations';
import { ActivatedRoute, Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { ConstraintEditDialogComponent } from '../constraint-edit-dialog/constraint-edit-dialog.component';
import { Constraint } from 'src/app/shared/models/constraint.model';
import { ValidationApiService } from 'src/app/core/services/validation-api.service';

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
  @Input() isRequired: boolean;
  @Output() constraintDisabledToggled: EventEmitter<boolean> = new EventEmitter();

  @HostBinding('class.constraint') hasConstraint = false;

  collapsed = true;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private dialog: MatDialog,
    private validationApiService: ValidationApiService
  ) {}

  ngOnInit(): void {
    const { enum: enumArray, regex, min, max, disabled } = this.constraint || {};
    this.hasConstraint =
      (this.constraint !== null && (enumArray || regex || min || max) !== null) || disabled;
  }

  openEditDialog(): void {
    const path = this.getPath();

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

  toggleConstraint(): void {
    const path = this.getPath();
    this.validationApiService
      .toggleConstraint(path)
      .subscribe(() => this.constraintDisabledToggled.emit(!this.constraint?.disabled));
  }

  private getPath(): string {
    return this.route.snapshot.url.map(segment => segment.path).join('/') + '/' + this.title;
  }
}
