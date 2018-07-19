import { Component, OnInit } from '@angular/core';
import { Certificado } from '../../models/certificado';
import { CertificadoService } from '../../services/certificado.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-certificado-edit',
  templateUrl: './certificado-edit.component.html',
  styleUrls: []
})
export class CertificadoEditComponent implements OnInit {

  certificado: Certificado = new Certificado();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private certificadoService: CertificadoService) {

  }

  actualizar(certificado: Certificado): void {
    this.certificadoService.update(certificado).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.certificadoService.getCertificado(params['id']))
      .subscribe(certificado => this.certificado = certificado);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}