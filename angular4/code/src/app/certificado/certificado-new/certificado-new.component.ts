import { Component, OnInit } from '@angular/core';
import { Certificado } from '../../models/certificado';
import { CertificadoService } from '../../services/certificado.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-certificado-new',
  templateUrl: './certificado-new.component.html',
  styleUrls: []
})
export class CertificadoNewComponent implements OnInit {

  certificado: Certificado;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private certificadoService: CertificadoService, private location: Location) { }

  ngOnInit() {
    this.certificado = new Certificado();
  }

  guardar(certificado: Certificado): void {

    this.certificadoService.create(certificado);
    this.display = true;

  }

  isNumber(n){

    if(n){
      if(isNaN(parseInt(n))){
        return false;
      }
      return true;
    }
    return false;
  }


  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}