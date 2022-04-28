import { NgModule} from '@angular/core';
import { SharedModule } from "../../shared/shared-module"
import { ReactiveFormsModule } from '@angular/forms';
import { PageRoutingModule } from './+routing';
import { WebServerService } from "../../server/web-server.service"
import { CommonModule } from '@angular/common';
import { HomepageComponent } from './homepage/homepage.component';
import {NgHttpLoaderModule} from "ng-http-loader";

const COMPONENTS = [HomepageComponent];

@NgModule({
  providers:[WebServerService],
  imports: [SharedModule,PageRoutingModule,ReactiveFormsModule,CommonModule,NgHttpLoaderModule],
  declarations: [...COMPONENTS]
})
export class HomeModule {}
