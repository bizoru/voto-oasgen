import { Component, OnInit } from '@angular/core';
import { EstamentoService } from '../../services/estamento.service';
import { Estamento } from '../../models/estamento';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-estamento',
  templateUrl: './estamento-view.component.html',
  styleUrls: []
})
export class EstamentoComponent implements OnInit {

  estamentos: Estamento[];
  estamento: Estamento;

  constructor(private estamentoService: EstamentoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.estamentoService.getEstamentos().then(estamentos => this.estamentos = estamentos);
  }

  newEstamento(): void {

    this.router.navigate(['/estamento/new']).then(() => null);
    this.globals.currentModule = 'Estamento';
  }

  editar(estamento: Estamento): void {
    this.estamento = estamento;
    this.router.navigate(['/estamento/edit', this.estamento._id ]);
  }

  borrar(estamento: Estamento): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar estamento?',
      accept: () => {
        this.estamentoService.delete(estamento._id)
          .then(response => this.estamentoService.getEstamentos().then(estamentos => this.estamentos = estamentos));
      }
    });
  }
}