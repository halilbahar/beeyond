import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConstraintControlComponent } from './constraint-control.component';

describe('ConstraintControlComponent', () => {
  let component: ConstraintControlComponent;
  let fixture: ComponentFixture<ConstraintControlComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConstraintControlComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ConstraintControlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
