import { Component, OnInit } from '@angular/core';
import { Configuracion } from '../../models/configuracion';
import { ConfiguracionService } from '../../services/configuracion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-configuracion-new',
  templateUrl: './configuracion-new.component.html',
  styleUrls: []
})
export class ConfiguracionNewComponent implements OnInit {

  configuracion: Configuracion;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private configuracionService: ConfiguracionService, private location: Location) { }

  ngOnInit() {
    this.configuracion = new Configuracion();
  }

  guardar(configuracion: Configuracion): void {

    this.configuracionService.create(configuracion);
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