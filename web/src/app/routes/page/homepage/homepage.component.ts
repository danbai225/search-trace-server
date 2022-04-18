import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { faWheelchair,faSearch } from '@fortawesome/free-solid-svg-icons';
import { NzMessageService } from 'ng-zorro-antd/message';
import { WebServerService } from "../../../server/web-server.service";
@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  data!: Array<string>;
  faWheelchair = faWheelchair;
  faSearch = faSearch;
  form: FormGroup;
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
 async handleSearchChange($event:string){
   let result : any;
   result = await this.server.getRequestpoo($event)
   if(result.msg === 'ok'){
    this.data = result.data
    console.log(this.data)
   }else{
     console.log(result.code);
   }
  }

}
