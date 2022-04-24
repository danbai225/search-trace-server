
import { Injectable } from "@angular/core";
import { Observable, Subject } from 'rxjs';

@Injectable()
export class LoadingService {

  private loadingSubject = new Subject<boolean>();

  constructor() {}

  getLoding(): Observable<boolean> {
    return this.loadingSubject.asObservable();
  }

  loading(showLoading: boolean) {
    this.loadingSubject.next(showLoading);
  }
}
