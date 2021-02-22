import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConstraintEditDialogComponent } from './constraint-edit-dialog.component';

describe('ConstraintEditDialogComponent', () => {
  let component: ConstraintEditDialogComponent;
  let fixture: ComponentFixture<ConstraintEditDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConstraintEditDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ConstraintEditDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
