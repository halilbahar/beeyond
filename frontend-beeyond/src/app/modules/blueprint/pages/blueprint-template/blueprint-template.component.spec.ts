import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BlueprintTemplateComponent } from './blueprint-template.component';

describe('BlueprintTemplateComponent', () => {
  let component: BlueprintTemplateComponent;
  let fixture: ComponentFixture<BlueprintTemplateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BlueprintTemplateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BlueprintTemplateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
