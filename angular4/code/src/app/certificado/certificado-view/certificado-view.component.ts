import { Component, OnInit } from '@angular/core';
import { CertificadoService } from '../../services/certificado.service';
import { Certificado } from '../../models/certificado';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-certificado',
  templateUrl: './certificado-view.component.html',
  styleUrls: []
})
export class CertificadoComponent implements OnInit {

  certificados: Certificado[];
  certificado: Certificado;

  constructor(private certificadoService: CertificadoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.certificadoService.getCertificados().then(certificados => this.certificados = certificados);
  }

  newCertificado(): void {

    this.router.navigate(['/certificado/new']).then(() => null);
    this.globals.currentModule = 'Certificado';
  }

  editar(certificado: Certificado): void {
    this.certificado = certificado;
    this.router.navigate(['/certificado/edit', this.certificado._id ]);
  }

  borrar(certificado: Certificado): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar certificado?',
      accept: () => {
        this.certificadoService.delete(certificado._id)
          .then(response => this.certificadoService.getCertificados().then(certificados => this.certificados = certificados));
      }
    });
  }
}