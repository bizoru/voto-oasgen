import { Component, OnInit } from '@angular/core';
import { ConfiguracionService } from '../../services/configuracion.service';
import { Configuracion } from '../../models/configuracion';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-configuracion',
  templateUrl: './configuracion-view.component.html',
  styleUrls: []
})
export class ConfiguracionComponent implements OnInit {

  configuracions: Configuracion[];
  configuracion: Configuracion;

  constructor(private configuracionService: ConfiguracionService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.configuracionService.getConfiguracions().then(configuracions => this.configuracions = configuracions);
  }

  newConfiguracion(): void {

    this.router.navigate(['/configuracion/new']).then(() => null);
    this.globals.currentModule = 'Configuracion';
  }

  editar(configuracion: Configuracion): void {
    this.configuracion = configuracion;
    this.router.navigate(['/configuracion/edit', this.configuracion._id ]);
  }

  borrar(configuracion: Configuracion): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar configuracion?',
      accept: () => {
        this.configuracionService.delete(configuracion._id)
          .then(response => this.configuracionService.getConfiguracions().then(configuracions => this.configuracions = configuracions));
      }
    });
  }
}