import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SHARED_ZORRO_MODULES } from "./shared-zorro.module";
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { sharedIconModule } from "./shared-nzicon";

@NgModule({
  declarations: [],
  imports: [ CommonModule,FontAwesomeModule,SHARED_ZORRO_MODULES,sharedIconModule],
  exports: [...SHARED_ZORRO_MODULES,FontAwesomeModule,sharedIconModule],
  providers: [],
})
export class SharedModule {}
