import { Component, OnInit } from '@angular/core';
import { EventEmitter, Input, Output } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { FormsModule } from '@angular/forms';
@Component({
  selector: 'app-premium',
  templateUrl: './premium.component.html',
  styleUrls: ['./premium.component.css']
})
export class PremiumComponent implements OnInit
{
  constructor(private http: HttpClient) { }
  ngOnInit()
  {
    const headers = new HttpHeaders({
      'content-type': 'application/json',
      'authorization': 'Bearer ' + localStorage.token
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22' + localStorage.orgName + '%22%5D', { headers }).subscribe((res) =>
    {
    });
  }
  premium = {}
  risk = ''
  mytoken_count = 10;
  required_token = 100;
  isgenerated = 'no';
  status = 'Generate Premium';
  changestatus()
  {
    switch (this.isgenerated)
    {
      case 'no':
        this.isgenerated = 'pending';
        this.status = 'Generating Premium';
        const headers = new HttpHeaders({
          'content-type': 'application/json',
          'authorization': 'Bearer ' + localStorage.token
        });
        this.http.post("http://localhost:4000/channels/mychannel/chaincodes/mycc",
          {
            "peers": ["peer0.org1.example.com", "peer0.org2.example.com", "peer0.org3.example.com"],
            "fcn": "calculateRiscFactor",
            "args": [localStorage.orgName, "3", "4", "9", "9"]
          }, { headers })
          .subscribe(
            data =>
            {
              alert(data)
              console.log("POST Request is successful ", data);
            },
            error =>
            {
              console.log("Error", error);
            }
          );
        break;
      case 'pending':
        this.isgenerated = 'yes';
        this.status = 'Please Wait...';
        break;
      default:
        break;
    }
  }
  @Output() callparent = new EventEmitter<boolean>();
  changecontent(x)
  {
    this.callparent.emit(x);
  }
}