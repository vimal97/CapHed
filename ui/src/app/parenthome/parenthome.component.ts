import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
@Component({
  selector: 'app-parenthome',
  templateUrl: './parenthome.component.html',
  styleUrls: ['./parenthome.component.css']
})
export class ParenthomeComponent implements OnInit
{
  userdata
  username
  policydata
  user1
  user2
  paymentuser1: string;
  paymentuser2: string;
  constructor(private http: HttpClient) { }
  ngOnInit()
  {
    this.username = localStorage.user;
    const headers = new HttpHeaders({
      'content-type': 'application/json',
      'authorization': 'Bearer ' + localStorage.token
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22' + localStorage.username + '%22%5D', { headers }).subscribe((res) =>
    {
      console.log(res);
      this.userdata = res;
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22policy1%22%5D', { headers }).subscribe((res) =>
    {
      console.log(res);
      this.policydata = res;
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22Child-1%22%5D', { headers }).subscribe((res) =>
    {
      console.log(res);
      this.user1 = res;
      if(this.user1.policytokens > 0)
      {
        this.paymentuser1 = "Paid";
      }
      else  
        this.paymentuser1 = "Unpaid";
      if(this.user1.policytokens > 0)
      {
        this.paymentuser2 = "Paid";
      }
      else
        this.paymentuser2 = "Unpaid";
    });
    this.http.get('http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22Child-2%22%5D', { headers }).subscribe((res) =>
    {
      console.log(res);
      this.user2 = res;
    });
  }
}