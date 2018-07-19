import { Component, OnInit } from '@angular/core';
import { Proceso } from '../../models/proceso';
import { ProcesoService } from '../../services/proceso.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-proceso-new',
  templateUrl: './proceso-new.component.html',
  styleUrls: []
})
export class ProcesoNewComponent implements OnInit {

  proceso: Proceso;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private procesoService: ProcesoService, private location: Location) { }

  ngOnInit() {
    this.proceso = new Proceso();
  }

  guardar(proceso: Proceso): void {

    this.procesoService.create(proceso);
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