import { Component, inject, OnInit, signal } from '@angular/core';
import { Equip } from '../types/model.type';
import { EquipService } from '../services/equip.service';

@Component({
  selector: 'app-equipment',
  standalone: true,
  imports: [],
  templateUrl: './equipment.component.html',
  styleUrl: './equipment.component.css'
})
export class EquipmentComponent implements OnInit {
  // equipsService = inject(EquipService)
  equips = signal<Array<Equip>>([])

  async ngOnInit(): Promise<void> {
    // this.equipsService.getEquip().subscribe((equips) => {
    //   this.equips.set(equips)
    // })
    const response = await fetch('http://localhost:8080/posts', {
      mode: "cors"
    })
    
    const json: Array<Equip> = await response.json();
    this.equips.set(json)
    console.log(json);
  }
}
