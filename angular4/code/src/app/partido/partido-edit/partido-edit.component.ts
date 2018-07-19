import { Component, OnInit } from '@angular/core';
import { Partido } from '../../models/partido';
import { PartidoService } from '../../services/partido.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-partido-edit',
  templateUrl: './partido-edit.component.html',
  styleUrls: []
})
export class PartidoEditComponent implements OnInit {

  partido: Partido = new Partido();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private partidoService: PartidoService) {

  }

  actualizar(partido: Partido): void {
    this.partidoService.update(partido).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.partidoService.getPartido(params['id']))
      .subscribe(partido => this.partido = partido);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}