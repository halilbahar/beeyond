import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApplicationContentComponent } from './application-content.component';

describe('ApplicationContentComponent', () => {
  let component: ApplicationContentComponent;
  let fixture: ComponentFixture<ApplicationContentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApplicationContentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApplicationContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
