import { Component, Input, OnInit } from '@angular/core';
import { AUTO_STYLE, animate, state, style, transition, trigger } from '@angular/animations';
import { Schema } from 'src/app/shared/models/schema.model';
import { ActivatedRoute, Router } from '@angular/router';

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
  @Input() hasRef: boolean;

  collapsed = true;

  constructor(private router: Router, private route: ActivatedRoute) {}

  ngOnInit(): void {}

  getGroupKindVersionName(schema: Schema) {
    const { group, kind, version } = schema['x-kubernetes-group-version-kind'][0];

    let groupString = '';
    if (group !== '') {
      groupString = '-' + group;
    }

    return kind + groupString + '-' + version;
  }

  navigate(path: string) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate([path], { relativeTo: this.route });
  }
}
