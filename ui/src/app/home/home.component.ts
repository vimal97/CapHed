import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  username
  userdata
  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.username = localStorage.user;
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

}
