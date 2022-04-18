import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import {  Router } from "@angular/router"
import { faCoffee,faUser } from '@fortawesome/free-solid-svg-icons';
import { WebServerService } from "../../server/web-server.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  faCoffee = faCoffee;
  validateForm!: FormGroup;
  user = faUser;
  users:string = ''
  constructor(private Server: WebServerService,private fb: FormBuilder, private router:Router) {}

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userName: [null,Validators.required],
      password: [null,Validators.required],
      remember: [true]
    });
  }
 async submitForm(){
    if (this.validateForm.valid) {
      let name = this.validateForm.value.userName;
      let pass = this.validateForm.value.password;
      let result : any;
      // token
      result = await this.Server.getLogin(name, pass);
      if(result.msg === 'ok'){
         this.router.navigate(['home']);
         localStorage.setItem('_token',result.data);
      }else{
        console.log('密码错误')
      }
    } else {
      Object.values(this.validateForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
}
