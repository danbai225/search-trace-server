import { Component, OnInit } from '@angular/core';
import { LoadingService } from '../../server/loading.service';

@Component({
  selector: 'app-loading',
  templateUrl: './loading.component.html',
  styleUrls: ['./loading.component.scss']
})
export class  LoadingComponent implements OnInit {

  public showLoading: boolean = false;//是否显示loading
  constructor(private loadingService: LoadingService) {}

  ngOnInit() {
    this.loadingService.getLoding().subscribe(loading => {
        this.showLoading = loading;
    });
  }
}
