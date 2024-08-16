import { Component } from '@angular/core';
import { AuthService } from '../../auth/auth.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent {
  loginForm: FormGroup;
  registerForm: FormGroup;

  constructor(private authService: AuthService, private fb: FormBuilder) {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });

    this.registerForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  login(): void {
    const { username, password } = this.loginForm.value;
    this.authService.login(username, password).subscribe();
  }

  register(): void {
    const { username, password } = this.registerForm.value;
    this.authService.register(username, password).subscribe();
}
}