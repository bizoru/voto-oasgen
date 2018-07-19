import { Component, OnInit } from '@angular/core';
import { Estamento } from '../../models/estamento';
import { EstamentoService } from '../../services/estamento.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-estamento-new',
  templateUrl: './estamento-new.component.html',
  styleUrls: []
})
export class EstamentoNewComponent implements OnInit {

  estamento: Estamento;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private estamentoService: EstamentoService, private location: Location) { }

  ngOnInit() {
    this.estamento = new Estamento();
  }

  guardar(estamento: Estamento): void {

    this.estamentoService.create(estamento);
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