// camera.component.ts
import { Component, ViewChild, ElementRef, OnInit, OnDestroy } from '@angular/core';

@Component({
  selector: 'app-camera',
  templateUrl: './camera.component.html',
  styleUrls: ['./camera.component.css']
})
export class CameraComponent implements OnInit, OnDestroy {
  @ViewChild('videoElement', { static: true }) videoElement!: ElementRef<HTMLVideoElement>;
  stream!: MediaStream;

  ngOnInit() {}

  // Start the camera
  async startCamera() {
    try {
      // Request access to the camera
      this.stream = await navigator.mediaDevices.getUserMedia({ video: true });
      this.videoElement.nativeElement.srcObject = this.stream;
    } catch (error) {
      console.error("Camera access error:", error);
    }
  }

  // Stop the camera when leaving the component
  ngOnDestroy() {
    if (this.stream) {
      this.stream.getTracks().forEach(track => track.stop());
    }
  }
}
