import { Component, OnInit } from '@angular/core';
import { Sufragante } from '../../models/sufragante';
import { SufraganteService } from '../../services/sufragante.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-sufragante-edit',
  templateUrl: './sufragante-edit.component.html',
  styleUrls: []
})
export class SufraganteEditComponent implements OnInit {

  sufragante: Sufragante = new Sufragante();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private sufraganteService: SufraganteService) {

  }

  actualizar(sufragante: Sufragante): void {
    this.sufraganteService.update(sufragante).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.sufraganteService.getSufragante(params['id']))
      .subscribe(sufragante => this.sufragante = sufragante);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}