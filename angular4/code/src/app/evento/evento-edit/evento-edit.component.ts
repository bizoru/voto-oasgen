import { Component, OnInit } from '@angular/core';
import { Evento } from '../../models/evento';
import { EventoService } from '../../services/evento.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-evento-edit',
  templateUrl: './evento-edit.component.html',
  styleUrls: []
})
export class EventoEditComponent implements OnInit {

  evento: Evento = new Evento();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private eventoService: EventoService) {

  }

  actualizar(evento: Evento): void {
    this.eventoService.update(evento).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.eventoService.getEvento(params['id']))
      .subscribe(evento => this.evento = evento);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}