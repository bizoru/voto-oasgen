import { Component, OnInit } from '@angular/core';
import { Voto } from '../../models/voto';
import { VotoService } from '../../services/voto.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-voto-new',
  templateUrl: './voto-new.component.html',
  styleUrls: []
})
export class VotoNewComponent implements OnInit {

  voto: Voto;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private votoService: VotoService, private location: Location) { }

  ngOnInit() {
    this.voto = new Voto();
  }

  guardar(voto: Voto): void {

    this.votoService.create(voto);
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