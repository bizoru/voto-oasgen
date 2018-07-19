import { Component, OnInit } from '@angular/core';
import { EventoService } from '../../services/evento.service';
import { Evento } from '../../models/evento';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-evento',
  templateUrl: './evento-view.component.html',
  styleUrls: []
})
export class EventoComponent implements OnInit {

  eventos: Evento[];
  evento: Evento;

  constructor(private eventoService: EventoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.eventoService.getEventos().then(eventos => this.eventos = eventos);
  }

  newEvento(): void {

    this.router.navigate(['/evento/new']).then(() => null);
    this.globals.currentModule = 'Evento';
  }

  editar(evento: Evento): void {
    this.evento = evento;
    this.router.navigate(['/evento/edit', this.evento._id ]);
  }

  borrar(evento: Evento): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar evento?',
      accept: () => {
        this.eventoService.delete(evento._id)
          .then(response => this.eventoService.getEventos().then(eventos => this.eventos = eventos));
      }
    });
  }
}