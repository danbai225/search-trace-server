import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router} from '@angular/router';
import { faWheelchair,faSearch,faTimes,faAngleDown,faAngleUp,faBars} from '@fortawesome/free-solid-svg-icons';
import { NumberInput } from 'ng-zorro-antd/core/types';
import { NzMessageService } from 'ng-zorro-antd/message';
import { WebServerService } from "../../../server/web-server.service";


@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  // 搜索关键字 data
  data!: Array<string>;
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
  isVisibleMiddle: boolean = false;

  // table select

  constructor(private formBuilder: FormBuilder,public msg: NzMessageService,private server:WebServerService,private router:Router) {
    this.form = this.formBuilder.group({
      comment: [null, [Validators.maxLength(100)]]
    });
  }
 async ngOnInit(): Promise<void> {
      // 判断 token
      let result:any;
      result = await this.server.getRequesInfo();
      if(result.msg !== 'ok'){
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
  }
  onCurrentPageDataChange($event:any){
    console.log($event);
  }
}
