import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';



import { LoadingComponent } from './loading/loading.component';


const COMPONENTS = [LoadingComponent];

@NgModule({

  imports: [CommonModule],
  declarations: [...COMPONENTS]
})
export class LoadingModule {}
