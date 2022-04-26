import { Injectable } from '@angular/core';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest} from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable()
export class Interceptor implements HttpInterceptor {
    intercept(req: HttpRequest<any>,next: HttpHandler): Observable<HttpEvent<any>> {
       const token :string | null = window.localStorage.getItem('_token');
        if (token) {
            // 设置请求头
            req = req.clone({
                setHeaders: {
                    'token': `${token}`,
                    'Content-Type': 'application/json'
                }
            });
        }
        return next.handle(req);
    }
}
