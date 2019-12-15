import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";

@Component({
  selector: 'app-buytoken',
  templateUrl: './buytoken.component.html',
  styleUrls: ['./buytoken.component.css']
})
export class BuytokenComponent implements OnInit
{
  userdata
  constructor(private http: HttpClient) { }
  ngOnInit()
  {
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
  selected = 'parent';
  pay()
  {
    window.location.href = '/payment'
  }
}