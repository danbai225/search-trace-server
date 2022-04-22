import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router} from '@angular/router';
import { faWheelchair,faSearch,faTimes,faAngleDown,faAngleUp,faBars } from '@fortawesome/free-solid-svg-icons';
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
  // 控制 搜索栏 展示
  isShow:boolean = true;
  // 判断 icon down uP
  isDownUp:boolean = true;
  faAngleDown = faAngleDown;
  faAngleUp = faAngleUp;
  form: FormGroup;
  $time:any;
  jion:string = '加入黑名单';
  out:string = '移除黑名单';
  searchString:string ='';
  //  分页 初始化
  PageIndex:string = '1';
  Total:string = '18';
  PageSize:string = '10';
  // 关键词
  word:string = '';
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
}
