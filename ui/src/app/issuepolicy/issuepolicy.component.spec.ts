import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { IssuepolicyComponent } from './issuepolicy.component';

describe('IssuepolicyComponent', () => {
  let component: IssuepolicyComponent;
  let fixture: ComponentFixture<IssuepolicyComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ IssuepolicyComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(IssuepolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
