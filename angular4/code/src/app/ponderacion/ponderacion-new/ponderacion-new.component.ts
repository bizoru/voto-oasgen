import { Component, OnInit } from '@angular/core';
import { Ponderacion } from '../../models/ponderacion';
import { PonderacionService } from '../../services/ponderacion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-ponderacion-new',
  templateUrl: './ponderacion-new.component.html',
  styleUrls: []
})
export class PonderacionNewComponent implements OnInit {

  ponderacion: Ponderacion;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private ponderacionService: PonderacionService, private location: Location) { }

  ngOnInit() {
    this.ponderacion = new Ponderacion();
  }

  guardar(ponderacion: Ponderacion): void {

    this.ponderacionService.create(ponderacion);
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