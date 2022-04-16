import { Component, OnInit } from '@angular/core';
import { faCoffee,faUser } from '@fortawesome/free-solid-svg-icons';
import { WebServerService } from "../../server/web-server.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  faCoffee = faCoffee;
  user = faUser;
  constructor(private Server: WebServerService) {}

  ngOnInit(): void {
    this.Server.getLogin();


  }
  handlelogin(){

  }
}
