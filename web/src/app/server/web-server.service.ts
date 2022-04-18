import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from "../../environments/environment";



@Injectable()
export class WebServerService {
  stieID: any;
  constructor (private http:HttpClient){}


  //  测试地址
   baseUrl(){
    return environment.apiUrl
  }
  // 请求登录
  getLogin(name:string,pass:string){
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/get_token`,{
        "name": name,
        "pass": pass
        }).subscribe(
        data => {
        resolve(data);
        });
    })
 }
//  请求词联想
  getRequestpoo(text:string){
     return new Promise((resolve,reject)=>{
      this.http.get(`${this.baseUrl()}/api/v1/word/associate/`,{
        params:{word:text}
        }).subscribe(
        data => {
        resolve(data);
        });
    })
     }
  }
