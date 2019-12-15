import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-topbar',
  templateUrl: './topbar.component.html',
  styleUrls: ['./topbar.component.css']
})
export class TopbarComponent implements OnInit {

  constructor() { }

  styleObject(): Object {
    if (localStorage.getItem("user") == "Parent1"){
        return {"background-color": "#57c34a"}
    }
    return {}
}

  ngOnInit() {
  }

}
