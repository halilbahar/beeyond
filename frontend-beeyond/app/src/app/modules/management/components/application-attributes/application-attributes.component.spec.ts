import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApplicationAttributesComponent } from './application-attributes.component';

describe('ApplicationAttributesComponent', () => {
  let component: ApplicationAttributesComponent;
  let fixture: ComponentFixture<ApplicationAttributesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApplicationAttributesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApplicationAttributesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
