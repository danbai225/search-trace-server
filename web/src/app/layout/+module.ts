import { NgModule } from '@angular/core';
import { SharedModule } from "../shared/shared-module"
import { PageRoutingModule } from './+routing';
import { LoginComponent } from './login/login.component';

const COMPONENTS = [LoginComponent];

@NgModule({
  imports: [SharedModule,PageRoutingModule],
  declarations: [...COMPONENTS]
})
export class LoginModule {}
