import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NamespaceDeleteDialogComponent } from './namespace-delete-dialog.component';

describe('NamespaceDeleteDialogComponent', () => {
  let component: NamespaceDeleteDialogComponent;
  let fixture: ComponentFixture<NamespaceDeleteDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NamespaceDeleteDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NamespaceDeleteDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
