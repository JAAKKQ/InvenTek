import { Component, NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CameraComponent } from './camera/camera.component';
import {WebcamModule} from 'ngx-webcam';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CameraComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})



export class AppComponent {
  title = 'InventTek';
}
