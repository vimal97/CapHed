import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";

@Component({
  selector: 'app-claim',
  templateUrl: './claim.component.html',
  styleUrls: ['./claim.component.css']
})
export class ClaimComponent implements OnInit {

  userdata
  constructor(private http: HttpClient) { }

  ngOnInit() {
    const headers = new HttpHeaders({
      'content-type': 'application/json',
      'authorization': 'Bearer ' + localStorage.token
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22'+localStorage.username+'%22%5D', { headers }).subscribe((res) =>
    {
      console.log(res);
      this.userdata = res;
    });
  }

  raise_claim(){
    const headers = new HttpHeaders({
      'content-type': 'application/json',
      'authorization': 'Bearer ' + localStorage.token
    });
    this.http.post("http://localhost:4000/channels/mychannel/chaincodes/mycc",
      {
        "peers": ["peer0.org1.example.com", "peer0.org2.example.com", "peer0.org3.example.com"],
        "fcn": "claim",
        "args": [localStorage.username, "3"]
      }, { headers })
      .subscribe(
        data =>
        {
          // alert(data)
          console.log("POST Request is successful ", data);
        },
        error =>
        {
          console.log("Error", error);
        }
      );
  }

}
