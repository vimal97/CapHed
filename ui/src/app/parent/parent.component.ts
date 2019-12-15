import { Component, OnInit } from '@angular/core';
@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrls: ['./parent.component.css']
})
export class ParentComponent implements OnInit
{
  constructor() { }
  ngOnInit()
  {
  }
  title = 'blockhash';
  content = 'issuepolicy'
  changecontent(x)
  {
    this.content = x;
  }
}