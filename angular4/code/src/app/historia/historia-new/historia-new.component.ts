import { Component, OnInit } from '@angular/core';
import { Historia } from '../../models/historia';
import { HistoriaService } from '../../services/historia.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-historia-new',
  templateUrl: './historia-new.component.html',
  styleUrls: []
})
export class HistoriaNewComponent implements OnInit {

  historia: Historia;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private historiaService: HistoriaService, private location: Location) { }

  ngOnInit() {
    this.historia = new Historia();
  }

  guardar(historia: Historia): void {

    this.historiaService.create(historia);
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