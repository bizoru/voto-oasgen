import { Component, OnInit } from '@angular/core';
import { CandidatoService } from '../../services/candidato.service';
import { Candidato } from '../../models/candidato';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-candidato',
  templateUrl: './candidato-view.component.html',
  styleUrls: []
})
export class CandidatoComponent implements OnInit {

  candidatos: Candidato[];
  candidato: Candidato;

  constructor(private candidatoService: CandidatoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.candidatoService.getCandidatos().then(candidatos => this.candidatos = candidatos);
  }

  newCandidato(): void {

    this.router.navigate(['/candidato/new']).then(() => null);
    this.globals.currentModule = 'Candidato';
  }

  editar(candidato: Candidato): void {
    this.candidato = candidato;
    this.router.navigate(['/candidato/edit', this.candidato._id ]);
  }

  borrar(candidato: Candidato): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar candidato?',
      accept: () => {
        this.candidatoService.delete(candidato._id)
          .then(response => this.candidatoService.getCandidatos().then(candidatos => this.candidatos = candidatos));
      }
    });
  }
}