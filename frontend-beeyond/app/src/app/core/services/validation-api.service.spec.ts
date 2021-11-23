import { TestBed } from '@angular/core/testing';

import { ValidationApiService } from './validation-api.service';

describe('ValidationApiService', () => {
  let service: ValidationApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ValidationApiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
