import { Component, OnInit } from '@angular/core';
import { VotoService } from '../../services/voto.service';
import { Voto } from '../../models/voto';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-voto',
  templateUrl: './voto-view.component.html',
  styleUrls: []
})
export class VotoComponent implements OnInit {

  votos: Voto[];
  voto: Voto;

  constructor(private votoService: VotoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.votoService.getVotos().then(votos => this.votos = votos);
  }

  newVoto(): void {

    this.router.navigate(['/voto/new']).then(() => null);
    this.globals.currentModule = 'Voto';
  }

  editar(voto: Voto): void {
    this.voto = voto;
    this.router.navigate(['/voto/edit', this.voto._id ]);
  }

  borrar(voto: Voto): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar voto?',
      accept: () => {
        this.votoService.delete(voto._id)
          .then(response => this.votoService.getVotos().then(votos => this.votos = votos));
      }
    });
  }
}