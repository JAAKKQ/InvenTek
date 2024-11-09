import {Component, OnInit} from '@angular/core';
import {WebcamModule} from 'ngx-webcam';

@Component({
  selector: 'app-camera',
  standalone: true,
  imports: [WebcamModule],
  templateUrl: './camera.component.html',
  styleUrl: './camera.component.css'
})

export class CameraComponent{

}
