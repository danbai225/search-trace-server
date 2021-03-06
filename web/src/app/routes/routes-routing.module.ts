import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
// layout


const routes: Routes = [
  { path: '', loadChildren: () => import('../layout/+module').then(m => m.LoginModule) },
  { path: '',loadChildren:() => import('../routes/page/+module').then(m => m.HomeModule)},
  // { path: 'exception', loadChildren: () => import('./exception/exception.module').then(m => m.ExceptionModule) },
  // 单页不包裹Layout
  // { path: '**', redirectTo: 'exception/404' }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {

    })
  ],
  exports: [RouterModule]
})
export class RouteRoutingModule {}
