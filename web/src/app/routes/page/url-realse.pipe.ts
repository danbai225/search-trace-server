import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'UrlRealsPipe'
})
export class UrlRealsPipePipe implements PipeTransform {



  transform(value: any, ...args: unknown[]): unknown {
    let string = value.match(/(\S*)\//)[1];
    return string;
  }

}
