import { Component, OnInit } from '@angular/core';
import { EventEmitter, Input, Output } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { FormsModule } from '@angular/forms';
@Component({
  selector: 'app-issuepolicy',
  templateUrl: './issuepolicy.component.html',
  styleUrls: ['./issuepolicy.component.css']
})
export class IssuepolicyComponent implements OnInit
{
  constructor(private http: HttpClient) { }
  ngOnInit()
  {
  }
  count = 500;
  issue()
  {
    const headers = new HttpHeaders({
      'content-type': 'application/json',
      'authorization': 'Bearer ' + localStorage.token
    });
    this.http.post("http://localhost:4000/channels/mychannel/chaincodes/mycc", {
      "peers": ["peer0.org1.example.com", "peer0.org2.example.com", "peer0.org3.example.com"],
      "fcn": "newPolicy",
      "args": [this.count, "10/12/2019", "12/12/2019"]
    }, { headers })
      .subscribe(
        data =>
        {
          console.log("POST Request is successful ", data);
          window.location.href = 'parent'
        },
        error =>
        {
          console.log("Error", error);
          alert("error")
        }
      );
  }
}