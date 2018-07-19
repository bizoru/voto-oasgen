import { Component, OnInit } from '@angular/core';
import { Usuario } from '../../models/usuario';
import { UsuarioService } from '../../services/usuario.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-usuario-new',
  templateUrl: './usuario-new.component.html',
  styleUrls: []
})
export class UsuarioNewComponent implements OnInit {

  usuario: Usuario;
  display = false;
  isNaN: Function = Number.isNaN;
  constructor(private usuarioService: UsuarioService, private location: Location) { }

  ngOnInit() {
    this.usuario = new Usuario();
  }

  guardar(usuario: Usuario): void {

    this.usuarioService.create(usuario);
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