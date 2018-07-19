import { Component, OnInit } from '@angular/core';
import { Votante } from '../../models/votante';
import { VotanteService } from '../../services/votante.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-votante-new',
  templateUrl: './votante-new.component.html',
  styleUrls: []
})
export class VotanteNewComponent implements OnInit {

  votante: Votante;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private votanteService: VotanteService, private location: Location) { }

  ngOnInit() {
    this.votante = new Votante();
  }

  guardar(votante: Votante): void {

    this.votanteService.create(votante);
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