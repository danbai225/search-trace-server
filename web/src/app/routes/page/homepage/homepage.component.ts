import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router} from '@angular/router';
import { faWheelchair,faSearch,faTimes,faAngleDown,faAngleUp,faBars} from '@fortawesome/free-solid-svg-icons';
import { NumberInput } from 'ng-zorro-antd/core/types';
import { NzMessageService } from 'ng-zorro-antd/message';
import { WebServerService } from "../../../server/web-server.service";
import {$e} from "@angular/compiler/src/chars";




@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  // 搜索关键字 data
  data!: Array<string>;
  datalength:number = 0;
  // 关键字搜索 lists
  lists:any;
  faWheelchair = faWheelchair;
  faSearch = faSearch;
  faTimes = faTimes;
  faBars = faBars;
  checked:boolean = true;
  // faSquare = faSquare;
  // faSquareCheck = faSquareCheck;
  // 控制 搜索栏 展示
  isShow:boolean = true;
  // 判断 icon down uP
  isDownUp:boolean = true;
  faAngleDown = faAngleDown;
  faAngleUp = faAngleUp;
  form: FormGroup;
  $time:any;
  searchString:string ='';
  //  分页 初始化
  PageIndex:string = '1';
  Total:string = '18';
  PageSize:string = '10';
  //  表格 分页
  tablePageIndex:string = '1';
  Totals:NumberInput = '10';
  tablePageSize:string = '5';
  nzPageSizeOptions:number[] =[5,10]
  start:number = 0;
  end:number = 5;
  // 关键词
  word:string = '';
  // 黑名单list
  blackList: any;
  blacklists:any;
  isVisibleTop:boolean = false;
  isVisible1:boolean = false;
  isVisibleMiddle: boolean = false;
  // 控制编辑
  isEdit:boolean = false;
  // 用户信息
  users:any;
  validateForm!: FormGroup;
  // 判断添加用户
  isUser:boolean = false;
  // 判断添加表
  visible2:boolean = false;
  // 规则名字
  Rulename:string = '';

  startEdit(id: string): void {
    console.log(1)
    this.isEdit = true;
  }

 async cancelEdit(id: string){
    console.log(id);
    let result : any;
     result = await this.server.getRuequestdellist(id);
     if(result.msg ==='ok'){
       console.log('删除成功')
       this.isEdit = false;
       this.isVisibleTop = false;
     }
  }
  // 修改
  saveEdit(id: string): void {
    console.log(id)
    this.isEdit = true;
  }

  updateEditCache(): void {

  }
  constructor(private formBuilder: FormBuilder,public msg: NzMessageService,private server:WebServerService,private router:Router,private fb: FormBuilder,private message: NzMessageService) {
    this.form = this.formBuilder.group({
      comment: [null, [Validators.maxLength(100)]]
    });
  }
  createMessage(type: string): void {
    this.message.create(type, `This is a message of ${type}`);
  }
 async ngOnInit(): Promise<void> {
   this.validateForm = this.fb.group({
     userName: [null, [Validators.required]],
     password: [null, [Validators.required]],
     email:[null,[Validators.required]],
     remember: [true]
   });
   const data = [];
   for (let i = 0; i < 100; i++) {
     data.push({
       id: `${i}`,
       name: `Edrward ${i}`,
       age: 32,
       address: `London Park no. ${i}`
     });
   }
   this.updateEditCache();
      // 判断 token
      let result:any;
      result = await this.server.getRequesInfo();
      if(result.msg !== 'ok'){
         this.createMessage(result.msg);
         this.router.navigate(['login']);
      }
  }
  handleSearch(){
    console.log('搜索')
  }
  // x input 清空
  handleError(){
    this.searchString = '';
  }
  // 退出
  handleQuit(){
    localStorage.setItem('token','');
    this.router.navigate(['login']);
  }

  // 关键字搜索
 async handleResule(word:string){
   this.word = word;
    let result:any;
    result = await this.server.getRequestKeyword(word);
    if( result.msg === 'ok'){
       this.lists = result.data;
       this.datalength = this.lists.list.length;
       this.isShow = false;
      //  分页数据
      this.PageIndex = this.lists.page_num.toString();
      this.Total = this.lists.page_total.toString();
      this.PageSize = this.lists.list.length.toString();
    }else{
      console.log(result);
    }
  }
  // 搜索关键字
 async handleSearchChange($event:any){
   this.antiShake($event,1000)
  }
  // 防抖
  antiShake($event:any,sleep:number){
   clearInterval(this.$time);
   this.$time = setTimeout(async () => {
   let result : any;
   result = await this.server.getRequestpoo($event)
   if(result.msg === 'ok'){
    this.data = result.data;
    this.datalength = this.data.length
   }else{
     console.log(result.code);
   }
    }, sleep);
  }
  // todo黑名单
  handleaddBlacklist(){
   console.log(1);
  }
  // 切换黑名单标识
  handleadIconDown(){
    this.isDownUp = !this.isDownUp;
  }
  //切换黑名单数据接送
  handleTodo(){
  }

  // 页面改变的回调
 async handlePageindexChange(page:number){
    let result:any;
    result = await this.server.getRequestKeyword(this.word,parseInt(this.PageSize),page);
    this.lists = result.data;
    this.isShow = false;
    this.PageIndex = this.lists.page_num.toString();
    this.Total = this.lists.page_total.toString();
    this.PageSize = this.lists.list.length.toString();
  }
  // 页数改变的回调
async  handlePageSize(pageSize:number){
    let result:any;
    result = await this.server.getRequestKeyword(this.word,pageSize,parseInt(this.PageIndex));
    console.log(result)
    this.lists = result.data;
    this.isShow = false;
    this.PageIndex = this.lists.page_num.toString();
    this.Total = this.lists.page_total.toString();
    this.PageSize = this.lists.list.length.toString();
  }
 // table 分页
 handletableindexChange(page:number){
   this.start = ((page-1) * parseInt(this.tablePageSize));
   this.end = page * parseInt(this.tablePageSize);
   this.changeblick(this.start,this.end);
 }
//  table 页数
handletableSize(pageSize:number){
  this.start = 0;
  this.end = pageSize;
  this.changeblick(this.start,this.end)

}
 async showModalTop(){
    this.isVisibleTop = true;
    let result:any;
    result = await  this.server.getRequesAddblacklist();
    if(result.msg === 'ok'){
      console.log(result)
      this.blacklists = result.data;
      console.log(this.blacklists);
      this.Totals = this.blacklists.length;
      this.changeblick(this.start,this.end);
    }else{
      console.log('表格数据请求错误')
    }

  }
  // 黑名单list 初始化 + 分页
  changeblick(start:number,end:number){
    this.blackList = this.blacklists.slice(start,end);
  }
  showModalMiddle(): void {
    this.isVisibleMiddle = true;
  }
  // checkbox事件
  handleCheckChange($event:any,data: any){

  }
  // 下拉框change mode 事件
  handleSelectChange($event:any,data: any){
  //  data.mode = $event;
  //  this.black.id = $event;
  //  this.black.mode = $event;

  }
  // 下拉框change match_pattern 事件
  handlePatternChange($event:any,data:any){
  //  data.match_pattern = $event;
  //  this.black.match_pattern = $event;
  }
  // 文本域change 事件
  handleChangemodel($event:any,data:any){
  // data.rules = $event;
  // this.black.rules = $event;
  }
 async handleOkTop() {
    let result:any;


      // result = await this.server.getRuequesblocklist(this.black.id,this.black.enable,this.black.id.mode,this.black.id.match_pattern,this.black.id.rules);
      // console.log(result);

  }

  handleCancelTop(): void {
    this.isVisibleTop = false;
    this.isEdit = false;
  }
  //用户管理
 async showUserAdmin(){
    this.isVisible1 = true;
    //    获取用户信息
    let resulruse : any;
    resulruse= await this.server.getRuequestuser();

    if(resulruse.msg == 'ok'){
      this.users = resulruse.data;
    }
    console.log(resulruse)
  }
  handleCancel(){
    this.isVisible1 = false;
  }
  handleOk(){
    this.isVisible1 = false;
    this.isUser = false;
  }
    //添加
  addRow(){
      this.visible2 = true;
  }
  closeDrawer(){
    console.log(1);
    this.visible2 = false;
  }
 async useconfirm(id:number){
    console.log(id)
       let result : any;
       result = await this.server.getRuequestedeluser(id);
       if(result.msg === 'ok'){
         this.isVisible1 = false;
         this.isUser = false;
       }
  }
 async handleaddUser(){
   this.isUser = true;

  }
 async submitForm() {
   //添加用户
   console.log('submit', this.validateForm.value);
   if(this.validateForm.value.remember){
     let rults: any;
     rults = await this.server.getRuequestedituser(this.validateForm.value.userName, this.validateForm.value.email, this.validateForm.value.password);
     console.log(rults)
   }

   // this.isUser = false;
 }
  handleAddrecord($event:any){
    console.log($event)
  }
  Adddomain($event:any){
    console.log($event)
  }
  onCurrentPageDataChange($event:any){
    console.log($event);
  }
}
