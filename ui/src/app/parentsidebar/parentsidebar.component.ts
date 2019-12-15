import { Component, OnInit, Output, EventEmitter } from '@angular/core';
@Component({
  selector: 'app-parentsidebar',
  templateUrl: './parentsidebar.component.html',
  styleUrls: ['./parentsidebar.component.css']
})
export class ParentsidebarComponent implements OnInit
{
  username
  constructor() { }
  ngOnInit()
  {
    this.username = localStorage.user
  }
  @Output() callparent = new EventEmitter<boolean>();
  changecontent(x)
  {
    this.callparent.emit(x);
  }
}