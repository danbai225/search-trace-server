import { NgModule } from '@angular/core';
import { SharedModule } from "../shared/shared-module"
import { ReactiveFormsModule } from '@angular/forms';
import { PageRoutingModule } from './+routing';
import { WebServerService } from "../server/web-server.service";

import { LoginComponent } from './login/login.component';
import {NgHttpLoaderModule} from "ng-http-loader";


const COMPONENTS = [LoginComponent];

@NgModule({
  providers:[WebServerService],
  imports: [SharedModule, PageRoutingModule, ReactiveFormsModule, NgHttpLoaderModule],
  declarations: [...COMPONENTS]
})
export class LoginModule {}
