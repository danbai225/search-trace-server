import { NgModule } from '@angular/core';
import { SharedModule } from "../shared/shared-module"
import { ReactiveFormsModule } from '@angular/forms';
import { PageRoutingModule } from './+routing';
import { WebServerService } from "../server/web-server.service";

import { LoginComponent } from './login/login.component';


const COMPONENTS = [LoginComponent];

@NgModule({
  providers:[WebServerService],
  imports: [SharedModule,PageRoutingModule, ReactiveFormsModule],
  declarations: [...COMPONENTS]
})
export class LoginModule {}
