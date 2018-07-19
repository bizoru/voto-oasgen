import { Component, OnInit } from '@angular/core';
import { Usuario } from '../../models/usuario';
import { UsuarioService } from '../../services/usuario.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-usuario-edit',
  templateUrl: './usuario-edit.component.html',
  styleUrls: []
})
export class UsuarioEditComponent implements OnInit {

  usuario: Usuario = new Usuario();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private usuarioService: UsuarioService) {

  }

  actualizar(usuario: Usuario): void {
    this.usuarioService.update(usuario).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.usuarioService.getUsuario(params['id']))
      .subscribe(usuario => this.usuario = usuario);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}