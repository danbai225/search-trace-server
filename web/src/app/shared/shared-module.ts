import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SHARED_ZORRO_MODULES } from "./shared-zorro.module";
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { sharedIconModule } from "./shared-nzicon";
import { FormsModule} from '@angular/forms';


// 拦截器
import { HTTP_INTERCEPTORS } from '@angular/common/http';
// 请求
import { Interceptor } from "../server/interceptors";
// 响应
import { CatchErrorInjector } from "../server/CatchErrorInjector";
const HttpInterceptorProviders = [
  { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  { provide: HTTP_INTERCEPTORS, useClass: CatchErrorInjector, multi: true },
]
@NgModule({
  declarations: [],
  imports: [ CommonModule,FontAwesomeModule,SHARED_ZORRO_MODULES,sharedIconModule,FormsModule],
  exports: [...SHARED_ZORRO_MODULES,FontAwesomeModule,sharedIconModule,FormsModule,CommonModule],
  providers: HttpInterceptorProviders,
})
export class SharedModule {}
