import { Component, OnInit } from '@angular/core';
import { SufraganteService } from '../../services/sufragante.service';
import { Sufragante } from '../../models/sufragante';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-sufragante',
  templateUrl: './sufragante-view.component.html',
  styleUrls: []
})
export class SufraganteComponent implements OnInit {

  sufragantes: Sufragante[];
  sufragante: Sufragante;

  constructor(private sufraganteService: SufraganteService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.sufraganteService.getSufragantes().then(sufragantes => this.sufragantes = sufragantes);
  }

  newSufragante(): void {

    this.router.navigate(['/sufragante/new']).then(() => null);
    this.globals.currentModule = 'Sufragante';
  }

  editar(sufragante: Sufragante): void {
    this.sufragante = sufragante;
    this.router.navigate(['/sufragante/edit', this.sufragante._id ]);
  }

  borrar(sufragante: Sufragante): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar sufragante?',
      accept: () => {
        this.sufraganteService.delete(sufragante._id)
          .then(response => this.sufraganteService.getSufragantes().then(sufragantes => this.sufragantes = sufragantes));
      }
    });
  }
}