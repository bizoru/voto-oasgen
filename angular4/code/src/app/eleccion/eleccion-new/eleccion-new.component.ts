import { Component, OnInit } from '@angular/core';
import { Eleccion } from '../../models/eleccion';
import { EleccionService } from '../../services/eleccion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-eleccion-new',
  templateUrl: './eleccion-new.component.html',
  styleUrls: []
})
export class EleccionNewComponent implements OnInit {

  eleccion: Eleccion;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private eleccionService: EleccionService, private location: Location) { }

  ngOnInit() {
    this.eleccion = new Eleccion();
  }

  guardar(eleccion: Eleccion): void {

    this.eleccionService.create(eleccion);
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