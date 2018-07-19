import { Component, OnInit } from '@angular/core';
import { PonderacionService } from '../../services/ponderacion.service';
import { Ponderacion } from '../../models/ponderacion';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-ponderacion',
  templateUrl: './ponderacion-view.component.html',
  styleUrls: []
})
export class PonderacionComponent implements OnInit {

  ponderacions: Ponderacion[];
  ponderacion: Ponderacion;

  constructor(private ponderacionService: PonderacionService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.ponderacionService.getPonderacions().then(ponderacions => this.ponderacions = ponderacions);
  }

  newPonderacion(): void {

    this.router.navigate(['/ponderacion/new']).then(() => null);
    this.globals.currentModule = 'Ponderacion';
  }

  editar(ponderacion: Ponderacion): void {
    this.ponderacion = ponderacion;
    this.router.navigate(['/ponderacion/edit', this.ponderacion._id ]);
  }

  borrar(ponderacion: Ponderacion): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar ponderacion?',
      accept: () => {
        this.ponderacionService.delete(ponderacion._id)
          .then(response => this.ponderacionService.getPonderacions().then(ponderacions => this.ponderacions = ponderacions));
      }
    });
  }
}