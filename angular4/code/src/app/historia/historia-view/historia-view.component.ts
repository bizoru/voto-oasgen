import { Component, OnInit } from '@angular/core';
import { HistoriaService } from '../../services/historia.service';
import { Historia } from '../../models/historia';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-historia',
  templateUrl: './historia-view.component.html',
  styleUrls: []
})
export class HistoriaComponent implements OnInit {

  historias: Historia[];
  historia: Historia;

  constructor(private historiaService: HistoriaService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.historiaService.getHistorias().then(historias => this.historias = historias);
  }

  newHistoria(): void {

    this.router.navigate(['/historia/new']).then(() => null);
    this.globals.currentModule = 'Historia';
  }

  editar(historia: Historia): void {
    this.historia = historia;
    this.router.navigate(['/historia/edit', this.historia._id ]);
  }

  borrar(historia: Historia): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar historia?',
      accept: () => {
        this.historiaService.delete(historia._id)
          .then(response => this.historiaService.getHistorias().then(historias => this.historias = historias));
      }
    });
  }
}