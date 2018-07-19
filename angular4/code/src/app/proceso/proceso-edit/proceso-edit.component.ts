import { Component, OnInit } from '@angular/core';
import { Proceso } from '../../models/proceso';
import { ProcesoService } from '../../services/proceso.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-proceso-edit',
  templateUrl: './proceso-edit.component.html',
  styleUrls: []
})
export class ProcesoEditComponent implements OnInit {

  proceso: Proceso = new Proceso();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private procesoService: ProcesoService) {

  }

  actualizar(proceso: Proceso): void {
    this.procesoService.update(proceso).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.procesoService.getProceso(params['id']))
      .subscribe(proceso => this.proceso = proceso);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}