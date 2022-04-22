import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { faWheelchair,faSearch,faTimes,faAngleDown,faAngleUp } from '@fortawesome/free-solid-svg-icons';
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
  // 判断 icon down uP
  isDownUp:boolean = true;
  faAngleDown = faAngleDown;
  faAngleUp = faAngleUp;
  form: FormGroup;
  $time:any;
  jion:string = '加入黑名单';
  out:string = '移除黑名单';
  searchString:string ='';
  constructor(private formBuilder: FormBuilder,public msg: NzMessageService,private server:WebServerService) {
    this.form = this.formBuilder.group({
      comment: [null, [Validators.maxLength(100)]]
    });
  }
  ngOnInit(): void {

  }
  handleSearch(){
    console.log('搜索')
  }
  // x input 清空
  handleError(){
    this.searchString = '';
  }

  // 关键字搜索
 async handleResule(word:string){
    let result:any;
    result = await this.server.getRequestKeyword(word);
    if( result.msg === 'ok'){
       this.lists = result.data;
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
    this.data = result.data
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
}
