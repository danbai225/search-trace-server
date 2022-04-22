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
  // 关键字搜索
  getRequestKeyword(keyWord:string,page_size:number=10,page_num:number=1){
    return new Promise((resolve,reject)=>{
      this.http.get(`${this.baseUrl()}/api/v1/trace/search_keyword`,{
        params:{key:keyWord,page_size,page_num}
        }).subscribe(
        data => {
        resolve(data);
        });
    })
  }
  }
