import { Component, OnInit } from '@angular/core';
import { Rol } from '../../models/rol';
import { RolService } from '../../services/rol.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-rol-edit',
  templateUrl: './rol-edit.component.html',
  styleUrls: []
})
export class RolEditComponent implements OnInit {

  rol: Rol = new Rol();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private rolService: RolService) {

  }

  actualizar(rol: Rol): void {
    this.rolService.update(rol).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.rolService.getRol(params['id']))
      .subscribe(rol => this.rol = rol);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}