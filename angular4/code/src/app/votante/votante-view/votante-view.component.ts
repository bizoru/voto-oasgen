import { Component, OnInit } from '@angular/core';
import { VotanteService } from '../../services/votante.service';
import { Votante } from '../../models/votante';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-votante',
  templateUrl: './votante-view.component.html',
  styleUrls: []
})
export class VotanteComponent implements OnInit {

  votantes: Votante[];
  votante: Votante;

  constructor(private votanteService: VotanteService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.votanteService.getVotantes().then(votantes => this.votantes = votantes);
  }

  newVotante(): void {

    this.router.navigate(['/votante/new']).then(() => null);
    this.globals.currentModule = 'Votante';
  }

  editar(votante: Votante): void {
    this.votante = votante;
    this.router.navigate(['/votante/edit', this.votante._id ]);
  }

  borrar(votante: Votante): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar votante?',
      accept: () => {
        this.votanteService.delete(votante._id)
          .then(response => this.votanteService.getVotantes().then(votantes => this.votantes = votantes));
      }
    });
  }
}