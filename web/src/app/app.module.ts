import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { RouteRoutingModule } from './routes/routes-routing.module';
import { AppComponent } from './app.component';
import { SharedModule } from "../app/shared/shared-module";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'; 




@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    RouteRoutingModule,
    SharedModule,
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {

}
