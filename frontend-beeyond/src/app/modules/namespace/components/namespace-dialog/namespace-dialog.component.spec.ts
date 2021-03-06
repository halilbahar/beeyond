import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NamespaceDialogComponent } from './namespace-dialog.component';

describe('NamespaceDialogComponent', () => {
  let component: NamespaceDialogComponent;
  let fixture: ComponentFixture<NamespaceDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NamespaceDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NamespaceDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
