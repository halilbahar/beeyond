import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TemplateApplicationPreviewDialogComponent } from './template-application-preview-dialog.component';

describe('TemplateApplicationPreviewDialogComponent', () => {
  let component: TemplateApplicationPreviewDialogComponent;
  let fixture: ComponentFixture<TemplateApplicationPreviewDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TemplateApplicationPreviewDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TemplateApplicationPreviewDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
