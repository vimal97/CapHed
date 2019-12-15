import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { PaymentComponent } from './payment/payment.component';
import { MainComponent } from './main/main.component';
import { ParentComponent } from './parent/parent.component';
const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'child', component: MainComponent },
  { path: 'payment', component: PaymentComponent },
  { path: 'parent', component: ParentComponent }];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }