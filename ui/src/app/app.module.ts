import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule } from '@angular/material';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { TopbarComponent } from './topbar/topbar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { PremiumComponent } from './premium/premium.component';
import { ClaimComponent } from './claim/claim.component';
import { BuytokenComponent } from './buytoken/buytoken.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { MainComponent } from './main/main.component';
import { PaymentComponent } from './payment/payment.component';
import { ParentComponent } from './parent/parent.component';
import { HttpClientModule } from '@angular/common/http';
import { ParentsidebarComponent } from './parentsidebar/parentsidebar.component';
import { ApproveclaimComponent } from './approveclaim/approveclaim.component';
import { IssuepolicyComponent } from './issuepolicy/issuepolicy.component';
import { ParenthomeComponent } from './parenthome/parenthome.component';
import { FormsModule } from '@angular/forms';
@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    TopbarComponent,
    PremiumComponent,
    ClaimComponent,
    BuytokenComponent,
    HomeComponent,
    LoginComponent,
    MainComponent,
    PaymentComponent,
    ParentComponent,
    ParentsidebarComponent,
    ApproveclaimComponent,
    IssuepolicyComponent,
    ParenthomeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatButtonModule,
    MatIconModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }