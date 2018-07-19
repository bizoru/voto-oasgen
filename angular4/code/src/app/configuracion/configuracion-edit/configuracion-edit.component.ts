import { Component, OnInit } from '@angular/core';
import { Configuracion } from '../../models/configuracion';
import { ConfiguracionService } from '../../services/configuracion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-configuracion-edit',
  templateUrl: './configuracion-edit.component.html',
  styleUrls: []
})
export class ConfiguracionEditComponent implements OnInit {

  configuracion: Configuracion = new Configuracion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private configuracionService: ConfiguracionService) {

  }

  actualizar(configuracion: Configuracion): void {
    this.configuracionService.update(configuracion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.configuracionService.getConfiguracion(params['id']))
      .subscribe(configuracion => this.configuracion = configuracion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}