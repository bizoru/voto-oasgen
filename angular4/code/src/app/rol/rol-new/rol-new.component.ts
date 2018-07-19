import { Component, OnInit } from '@angular/core';
import { Rol } from '../../models/rol';
import { RolService } from '../../services/rol.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-rol-new',
  templateUrl: './rol-new.component.html',
  styleUrls: []
})
export class RolNewComponent implements OnInit {

  rol: Rol;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private rolService: RolService, private location: Location) { }

  ngOnInit() {
    this.rol = new Rol();
  }

  guardar(rol: Rol): void {

    this.rolService.create(rol);
    this.display = true;

  }

  isNumber(n){

    if(n){
      if(isNaN(parseInt(n))){
        return false;
      }
      return true;
    }
    return false;
  }


  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}