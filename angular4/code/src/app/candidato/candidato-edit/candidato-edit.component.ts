import { Component, OnInit } from '@angular/core';
import { Candidato } from '../../models/candidato';
import { CandidatoService } from '../../services/candidato.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-candidato-edit',
  templateUrl: './candidato-edit.component.html',
  styleUrls: []
})
export class CandidatoEditComponent implements OnInit {

  candidato: Candidato = new Candidato();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private candidatoService: CandidatoService) {

  }

  actualizar(candidato: Candidato): void {
    this.candidatoService.update(candidato).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.candidatoService.getCandidato(params['id']))
      .subscribe(candidato => this.candidato = candidato);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}