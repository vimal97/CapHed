import { Component, OnInit } from '@angular/core';
import { EventEmitter, Input, Output } from '@angular/core';
@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit
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