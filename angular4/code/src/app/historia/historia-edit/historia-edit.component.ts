import { Component, OnInit } from '@angular/core';
import { Historia } from '../../models/historia';
import { HistoriaService } from '../../services/historia.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-historia-edit',
  templateUrl: './historia-edit.component.html',
  styleUrls: []
})
export class HistoriaEditComponent implements OnInit {

  historia: Historia = new Historia();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private historiaService: HistoriaService) {

  }

  actualizar(historia: Historia): void {
    this.historiaService.update(historia).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.historiaService.getHistoria(params['id']))
      .subscribe(historia => this.historia = historia);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}