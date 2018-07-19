import { Component, OnInit } from '@angular/core';
import { Variable } from '../../models/variable';
import { VariableService } from '../../services/variable.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-variable-new',
  templateUrl: './variable-new.component.html',
  styleUrls: []
})
export class VariableNewComponent implements OnInit {

  variable: Variable;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private variableService: VariableService, private location: Location) { }

  ngOnInit() {
    this.variable = new Variable();
  }

  guardar(variable: Variable): void {

    this.variableService.create(variable);
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