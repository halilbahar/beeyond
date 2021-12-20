import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BlueprintNewComponent } from './blueprint-new.component';

describe('BlueprintNewComponent', () => {
  let component: BlueprintNewComponent;
  let fixture: ComponentFixture<BlueprintNewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BlueprintNewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BlueprintNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
