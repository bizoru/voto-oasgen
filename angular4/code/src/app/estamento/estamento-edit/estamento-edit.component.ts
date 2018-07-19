import { Component, OnInit } from '@angular/core';
import { Estamento } from '../../models/estamento';
import { EstamentoService } from '../../services/estamento.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-estamento-edit',
  templateUrl: './estamento-edit.component.html',
  styleUrls: []
})
export class EstamentoEditComponent implements OnInit {

  estamento: Estamento = new Estamento();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private estamentoService: EstamentoService) {

  }

  actualizar(estamento: Estamento): void {
    this.estamentoService.update(estamento).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.estamentoService.getEstamento(params['id']))
      .subscribe(estamento => this.estamento = estamento);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}