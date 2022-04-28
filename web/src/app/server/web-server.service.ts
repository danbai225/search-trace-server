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
  // 获取用户信息
  getRequesInfo(){
    return new Promise((resolve,reject)=>{
      this.http.get(`${this.baseUrl()}/api/v1/user/info`).subscribe(
        data => {
        resolve(data);
        });
    })
  }
  // 获取黑名单信息
  getRequesAddblacklist(){
    return new Promise((resolve,reject)=>{
      this.http.get(`${this.baseUrl()}/api/v1/blacklist/list`).subscribe(
        data => {
        resolve(data);
        });
    })
  }
  //添加黑名单
  getRuequesaddblocklist(oneEnable:boolean,oneMode:number,oneMatch_pattern:number,Rules:string,Onename:string){
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/v1/blacklist/add`,{
        enable:oneEnable,
        mode:oneMode,
        match_pattern:oneMatch_pattern,
        rules:Rules,
        name:Onename
      }).subscribe(
        data => {
          resolve(data);
        });
    })
  }
  //修改黑名单
  getRuequesblocklist(oneId:number,oneEnable:boolean,oneMode:BigInteger,oneMatch_pattern:BigInteger,Rules:string){
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/v1/blacklist/add`,{
        id:oneId,
        enable:oneEnable,
        mode:oneMode,
        match_pattern:oneMatch_pattern,
        rules:Rules
      }).subscribe(
        data => {
        resolve(data);
        });
    })
  }
  //删除黑名单
  getRuequestdellist(oneId:string){
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/v1/blacklist/del`,{
        id:oneId,
      }).subscribe(
        data => {
          resolve(data);
        });
    })
  }
//  获取用户信息
  getRuequestuser(){
    return new Promise((resolve,reject)=>{
      this.http.get(`${this.baseUrl()}/api/v1/user/list`).subscribe(
        data => {
          resolve(data);
        });
    })
  }
// 添加编辑用户信息
  getRuequestedituser(Onename:string,Oneemail:string,Onepass:string){
    console.log(Onename,Oneemail,Onepass)
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/v1/user/add`,{
        name:Onename,
        email:Oneemail,
        pass:Onepass
      }).subscribe(
        data => {
          resolve(data);
        });
    })
  }
  //删除用户信息
  getRuequestedeluser(id:number,){
    return new Promise((resolve,reject)=>{
      this.http.post(`${this.baseUrl()}/api/v1/user/del`,{
        id:id
      }).subscribe(
        data => {
          resolve(data);
        });
    })
  }
}
