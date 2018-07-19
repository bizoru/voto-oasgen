import { Component, OnInit } from '@angular/core';
import { ProcesoService } from '../../services/proceso.service';
import { Proceso } from '../../models/proceso';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-proceso',
  templateUrl: './proceso-view.component.html',
  styleUrls: []
})
export class ProcesoComponent implements OnInit {

  procesos: Proceso[];
  proceso: Proceso;

  constructor(private procesoService: ProcesoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.procesoService.getProcesos().then(procesos => this.procesos = procesos);
  }

  newProceso(): void {

    this.router.navigate(['/proceso/new']).then(() => null);
    this.globals.currentModule = 'Proceso';
  }

  editar(proceso: Proceso): void {
    this.proceso = proceso;
    this.router.navigate(['/proceso/edit', this.proceso._id ]);
  }

  borrar(proceso: Proceso): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar proceso?',
      accept: () => {
        this.procesoService.delete(proceso._id)
          .then(response => this.procesoService.getProcesos().then(procesos => this.procesos = procesos));
      }
    });
  }
}