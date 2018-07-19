import { Component, OnInit } from '@angular/core';
import { TarjetonService } from '../../services/tarjeton.service';
import { Tarjeton } from '../../models/tarjeton';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-tarjeton',
  templateUrl: './tarjeton-view.component.html',
  styleUrls: []
})
export class TarjetonComponent implements OnInit {

  tarjetons: Tarjeton[];
  tarjeton: Tarjeton;

  constructor(private tarjetonService: TarjetonService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.tarjetonService.getTarjetons().then(tarjetons => this.tarjetons = tarjetons);
  }

  newTarjeton(): void {

    this.router.navigate(['/tarjeton/new']).then(() => null);
    this.globals.currentModule = 'Tarjeton';
  }

  editar(tarjeton: Tarjeton): void {
    this.tarjeton = tarjeton;
    this.router.navigate(['/tarjeton/edit', this.tarjeton._id ]);
  }

  borrar(tarjeton: Tarjeton): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar tarjeton?',
      accept: () => {
        this.tarjetonService.delete(tarjeton._id)
          .then(response => this.tarjetonService.getTarjetons().then(tarjetons => this.tarjetons = tarjetons));
      }
    });
  }
}