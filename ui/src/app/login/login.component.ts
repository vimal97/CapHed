import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username;
  password;
  constructor() { }

  login(){
    this.username = ((document.getElementById("login") as HTMLInputElement).value);
    this.password = ((document.getElementById("password") as HTMLInputElement).value);
    if(this.username == "Parent1" && this.password == "12345")
        {
            localStorage.setItem("user",this.username);
            localStorage.setItem("orgName","org1");
            localStorage.setItem("username","Parent-1");
            localStorage.setItem("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYwNDI2MjUsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NzYwMDY2MjV9.PO19C7otYl3YiuW8HgK2l0slR_ALZ9vqJScBABEA1s8");
            console.log(this.username);
            window.location.href = "http://localhost:4200/parent";
        }
        else if(this.username == "Child1" && this.password == "12345"){
            localStorage.setItem("user",this.username);
            localStorage.setItem("orgName","org2");
            localStorage.setItem("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYwNDI2MjUsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NzYwMDY2MjV9.PO19C7otYl3YiuW8HgK2l0slR_ALZ9vqJScBABEA1s8");
            localStorage.setItem("username","Child-1");
            window.location.href = "http://localhost:4200/child";
        }
        else if(this.username == "Child2" && this.password == "12345"){
            localStorage.setItem("user",this.username);
            localStorage.setItem("orgName","org3");
            localStorage.setItem("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYwNDI2MjUsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NzYwMDY2MjV9.PO19C7otYl3YiuW8HgK2l0slR_ALZ9vqJScBABEA1s8");
            localStorage.setItem("username","Child-2");
            window.location.href = "http://localhost:4200/child";
        }
        else{
            alert("Unknown user!");
        }
  }

  ngOnInit() {
  }

}
