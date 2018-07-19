import { Component, OnInit } from '@angular/core';
import { Sufragante } from '../../models/sufragante';
import { SufraganteService } from '../../services/sufragante.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-sufragante-new',
  templateUrl: './sufragante-new.component.html',
  styleUrls: []
})
export class SufraganteNewComponent implements OnInit {

  sufragante: Sufragante;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private sufraganteService: SufraganteService, private location: Location) { }

  ngOnInit() {
    this.sufragante = new Sufragante();
  }

  guardar(sufragante: Sufragante): void {

    this.sufraganteService.create(sufragante);
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