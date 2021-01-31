import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RootConstraintComponent } from './root-constraint.component';

describe('RootConstraintComponent', () => {
  let component: RootConstraintComponent;
  let fixture: ComponentFixture<RootConstraintComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RootConstraintComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RootConstraintComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
