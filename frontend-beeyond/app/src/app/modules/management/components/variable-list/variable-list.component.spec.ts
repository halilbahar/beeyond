import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { VariableListComponent } from './variable-list.component';

describe('TemplateListComponent', () => {
  let component: VariableListComponent;
  let fixture: ComponentFixture<VariableListComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [VariableListComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VariableListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
