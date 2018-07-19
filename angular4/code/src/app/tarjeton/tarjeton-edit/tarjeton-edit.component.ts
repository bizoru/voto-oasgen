import { Component, OnInit } from '@angular/core';
import { Tarjeton } from '../../models/tarjeton';
import { TarjetonService } from '../../services/tarjeton.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-tarjeton-edit',
  templateUrl: './tarjeton-edit.component.html',
  styleUrls: []
})
export class TarjetonEditComponent implements OnInit {

  tarjeton: Tarjeton = new Tarjeton();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private tarjetonService: TarjetonService) {

  }

  actualizar(tarjeton: Tarjeton): void {
    this.tarjetonService.update(tarjeton).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.tarjetonService.getTarjeton(params['id']))
      .subscribe(tarjeton => this.tarjeton = tarjeton);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}