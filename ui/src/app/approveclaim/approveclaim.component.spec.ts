import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ApproveclaimComponent } from './approveclaim.component';

describe('ApproveclaimComponent', () => {
  let component: ApproveclaimComponent;
  let fixture: ComponentFixture<ApproveclaimComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ApproveclaimComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApproveclaimComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
