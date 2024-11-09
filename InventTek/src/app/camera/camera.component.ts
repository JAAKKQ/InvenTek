import {Component, OnInit} from '@angular/core';
import { Subject, Observable } from 'rxjs';
import { WebcamModule, WebcamImage, WebcamInitError, WebcamUtil } from 'ngx-webcam';

@Component({
  selector: 'app-camera',
  standalone: true,
  imports: [WebcamModule],
  templateUrl: './camera.component.html',
  styleUrl: './camera.component.css'
})

export class CameraComponent{
  // Observable triggers for the snapshot and camera switch
  private trigger: Subject<void> = new Subject<void>();

  // Variable to store the captured image
  public webcamImage: WebcamImage | null = null;

  // Snapshot trigger as an observable
  public get triggerObservable(): Observable<void> {
    return this.trigger.asObservable();
  }

  // Trigger snapshot
  public takeSnapshot(): void {
    this.trigger.next();
  }

  // Handle the captured image
  public handleImage(webcamImage: WebcamImage): void {
    console.info('Received webcam image', webcamImage);
    this.webcamImage = webcamImage;
  }
}
