import { NgModule } from '@angular/core';
import { SharedModule } from "../shared/shared-module"
import { PageRoutingModule } from './+routing';
import { LoginComponent } from './login/login.component';
import { WebServerService } from "../server/web-server.service";
const COMPONENTS = [LoginComponent];

@NgModule({
  providers:[WebServerService],
  imports: [SharedModule,PageRoutingModule],
  declarations: [...COMPONENTS]
})
export class LoginModule {}
