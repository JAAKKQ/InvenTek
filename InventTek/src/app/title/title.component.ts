import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-title',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './title.component.html',
  styleUrl: './title.component.css'
})
export class TitleComponent {

}
