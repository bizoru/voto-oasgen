import { Component, OnInit } from '@angular/core';
import { Candidato } from '../../models/candidato';
import { CandidatoService } from '../../services/candidato.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-candidato-new',
  templateUrl: './candidato-new.component.html',
  styleUrls: []
})
export class CandidatoNewComponent implements OnInit {

  candidato: Candidato;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private candidatoService: CandidatoService, private location: Location) { }

  ngOnInit() {
    this.candidato = new Candidato();
  }

  guardar(candidato: Candidato): void {

    this.candidatoService.create(candidato);
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