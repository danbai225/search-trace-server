import { HttpInterceptor, HttpRequest, HttpHandler } from '@angular/common/http';
import { Injectable } from '@angular/core';
	import { Router ,NavigationExtras} from '@angular/router';
import { catchError, retry } from 'rxjs/operators';
@Injectable()
export class CatchErrorInjector implements HttpInterceptor {
  constructor(public router:Router){}
  intercept(req: HttpRequest<any>, next: HttpHandler) {
    return next.handle(req).pipe(
      catchError(err => {
        // token 判空处理
        // 页面响应 loadin 加载
        // 登录持久化
        // 页面骨架
        // 路由跳转
        if(err.error=== 'Unauthorized'){
          console.log('token过期')
          this.router.navigate(['login'])
        }
        throw err;
      }),
    );
  }
}
