import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConstraintAccordionComponent } from './constraint-accordion.component';

describe('ConstraintAccordionComponent', () => {
  let component: ConstraintAccordionComponent;
  let fixture: ComponentFixture<ConstraintAccordionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConstraintAccordionComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ConstraintAccordionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
