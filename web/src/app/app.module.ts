import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';


import { RouteRoutingModule } from './routes/routes-routing.module';
import { AppComponent } from './app.component';
import { SharedModule } from "../app/shared/shared-module";





@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    RouteRoutingModule,
    SharedModule,

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }