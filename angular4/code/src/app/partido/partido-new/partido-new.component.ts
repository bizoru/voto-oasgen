import { Component, OnInit } from '@angular/core';
import { Partido } from '../../models/partido';
import { PartidoService } from '../../services/partido.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-partido-new',
  templateUrl: './partido-new.component.html',
  styleUrls: []
})
export class PartidoNewComponent implements OnInit {

  partido: Partido;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private partidoService: PartidoService, private location: Location) { }

  ngOnInit() {
    this.partido = new Partido();
  }

  guardar(partido: Partido): void {

    this.partidoService.create(partido);
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