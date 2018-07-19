import { Component, OnInit } from '@angular/core';
import { PartidoService } from '../../services/partido.service';
import { Partido } from '../../models/partido';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-partido',
  templateUrl: './partido-view.component.html',
  styleUrls: []
})
export class PartidoComponent implements OnInit {

  partidos: Partido[];
  partido: Partido;

  constructor(private partidoService: PartidoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.partidoService.getPartidos().then(partidos => this.partidos = partidos);
  }

  newPartido(): void {

    this.router.navigate(['/partido/new']).then(() => null);
    this.globals.currentModule = 'Partido';
  }

  editar(partido: Partido): void {
    this.partido = partido;
    this.router.navigate(['/partido/edit', this.partido._id ]);
  }

  borrar(partido: Partido): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar partido?',
      accept: () => {
        this.partidoService.delete(partido._id)
          .then(response => this.partidoService.getPartidos().then(partidos => this.partidos = partidos));
      }
    });
  }
}