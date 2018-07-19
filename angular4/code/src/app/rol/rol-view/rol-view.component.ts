import { Component, OnInit } from '@angular/core';
import { RolService } from '../../services/rol.service';
import { Rol } from '../../models/rol';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-rol',
  templateUrl: './rol-view.component.html',
  styleUrls: []
})
export class RolComponent implements OnInit {

  rols: Rol[];
  rol: Rol;

  constructor(private rolService: RolService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.rolService.getRols().then(rols => this.rols = rols);
  }

  newRol(): void {

    this.router.navigate(['/rol/new']).then(() => null);
    this.globals.currentModule = 'Rol';
  }

  editar(rol: Rol): void {
    this.rol = rol;
    this.router.navigate(['/rol/edit', this.rol._id ]);
  }

  borrar(rol: Rol): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar rol?',
      accept: () => {
        this.rolService.delete(rol._id)
          .then(response => this.rolService.getRols().then(rols => this.rols = rols));
      }
    });
  }
}