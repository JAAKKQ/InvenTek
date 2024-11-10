import { Routes } from '@angular/router';

export const routes: Routes = [{
    path: '',
    pathMatch: 'full',
    loadComponent: () => {
        return import('./title/title.component').then((m) => m.TitleComponent)
    },
},{
    path: 'dashboard',
    pathMatch: 'full',
    loadComponent: () => {
        return import('./dashboard/dashboard.component').then((m) => m.DashboardComponent)
    },
},{
    path: 'equipment',
    pathMatch: 'full',
    loadComponent: () => {
        return import('./equipment/equipment.component').then((m) => m.EquipmentComponent)
    },
},{
    path: 'equipment-detail',
    pathMatch: 'full',
    loadComponent: () => {
        return import('./equipment-detail/equipment-detail.component').then((m) => m.EquipmentDetailComponent)
    },
},{
    path: 'camera',
    pathMatch: 'full',
    loadComponent: () => {
        return import('./camera/camera.component').then((m) => m.CameraComponent)
    },
}];
