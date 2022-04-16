import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { IconDefinition } from '@ant-design/icons-angular';
import {
  LockOutline,
  UserOutline
} from '@ant-design/icons-angular/icons';

const icons: IconDefinition[] = [
  UserOutline,
  LockOutline
];
@NgModule({
  declarations: [],
  imports: [ CommonModule, NzIconModule.forChild(icons),],
  exports: [],
  providers: [],
})
export class sharedIconModule {}
