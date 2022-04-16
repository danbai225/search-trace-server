import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from "../../environments/environment";

@Injectable()
export class WebServerService {
  stieID: any;

  constructor(private http: HttpClient) { }
  //  测试地址
   baseUrl(){
    return environment.apiUrl
  }
  // 请求登录
  getLogin(){
    console.log(`${this.baseUrl()}/api/get_token`)
    this.http.post(`${this.baseUrl()}/api/get_token`,{
      "name": "admin",
      "pass": "123456"
      }).subscribe(
      data => {
      console.log("DELETE Request is successful ", data);
      });
  }

}
