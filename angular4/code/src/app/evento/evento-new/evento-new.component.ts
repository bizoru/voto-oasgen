import { Component, OnInit } from '@angular/core';
import { Evento } from '../../models/evento';
import { EventoService } from '../../services/evento.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-evento-new',
  templateUrl: './evento-new.component.html',
  styleUrls: []
})
export class EventoNewComponent implements OnInit {

  evento: Evento;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private eventoService: EventoService, private location: Location) { }

  ngOnInit() {
    this.evento = new Evento();
  }

  guardar(evento: Evento): void {

    this.eventoService.create(evento);
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