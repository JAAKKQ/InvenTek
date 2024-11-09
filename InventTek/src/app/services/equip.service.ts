import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Equip } from '../types/model.type';

@Injectable({
  providedIn: 'root'
})
export class EquipService {
  url = 'http://localhost:8080'
  http = inject(HttpClient)
  getEquip(): Observable<any> {
    return this.http.get<Array<Equip>>(this.url);
  }
}
