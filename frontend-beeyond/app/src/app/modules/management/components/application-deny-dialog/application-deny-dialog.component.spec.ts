import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApplicationDenyDialogComponent } from './application-deny-dialog.component';

describe('ApplicationDenyDialogComponent', () => {
  let component: ApplicationDenyDialogComponent;
  let fixture: ComponentFixture<ApplicationDenyDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApplicationDenyDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApplicationDenyDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
