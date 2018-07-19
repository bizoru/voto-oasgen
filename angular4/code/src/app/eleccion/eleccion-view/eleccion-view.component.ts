import { Component, OnInit } from '@angular/core';
import { EleccionService } from '../../services/eleccion.service';
import { Eleccion } from '../../models/eleccion';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-eleccion',
  templateUrl: './eleccion-view.component.html',
  styleUrls: []
})
export class EleccionComponent implements OnInit {

  eleccions: Eleccion[];
  eleccion: Eleccion;

  constructor(private eleccionService: EleccionService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.eleccionService.getEleccions().then(eleccions => this.eleccions = eleccions);
  }

  newEleccion(): void {

    this.router.navigate(['/eleccion/new']).then(() => null);
    this.globals.currentModule = 'Eleccion';
  }

  editar(eleccion: Eleccion): void {
    this.eleccion = eleccion;
    this.router.navigate(['/eleccion/edit', this.eleccion._id ]);
  }

  borrar(eleccion: Eleccion): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar eleccion?',
      accept: () => {
        this.eleccionService.delete(eleccion._id)
          .then(response => this.eleccionService.getEleccions().then(eleccions => this.eleccions = eleccions));
      }
    });
  }
}