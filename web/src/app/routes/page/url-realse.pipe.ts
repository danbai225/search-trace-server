import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'UrlRealsPipe'
})
export class UrlRealsPipePipe implements PipeTransform {

  transform(value: any, ...args: unknown[]): unknown {
    let Url = value.indexOf('/');
    let string = value.indexOf('/',Url+ 2);
    return  value.slice(0,string);
  }

}
