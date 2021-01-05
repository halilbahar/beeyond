import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApplicationPreviewDialogComponent } from './application-preview-dialog.component';

describe('ApplicationPreviewDialogComponent', () => {
  let component: ApplicationPreviewDialogComponent;
  let fixture: ComponentFixture<ApplicationPreviewDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApplicationPreviewDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApplicationPreviewDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
