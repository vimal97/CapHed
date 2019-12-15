import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ParentsidebarComponent } from './parentsidebar.component';

describe('ParentsidebarComponent', () => {
  let component: ParentsidebarComponent;
  let fixture: ComponentFixture<ParentsidebarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ParentsidebarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ParentsidebarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
