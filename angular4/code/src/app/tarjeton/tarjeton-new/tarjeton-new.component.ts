import { Component, OnInit } from '@angular/core';
import { Tarjeton } from '../../models/tarjeton';
import { TarjetonService } from '../../services/tarjeton.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-tarjeton-new',
  templateUrl: './tarjeton-new.component.html',
  styleUrls: []
})
export class TarjetonNewComponent implements OnInit {

  tarjeton: Tarjeton;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private tarjetonService: TarjetonService, private location: Location) { }

  ngOnInit() {
    this.tarjeton = new Tarjeton();
  }

  guardar(tarjeton: Tarjeton): void {

    this.tarjetonService.create(tarjeton);
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