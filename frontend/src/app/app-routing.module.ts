import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SongListComponent } from './components/songlist/songlist.component';
import { UserComponent } from './components/user/user.component';
import { AuthGuard } from './auth/auth.guard';

const routes: Routes = [
  { path: 'login', component: UserComponent },
  { path: 'songlists', component: SongListComponent, canActivate: [AuthGuard] },
  { path: '', redirectTo: '/login', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
