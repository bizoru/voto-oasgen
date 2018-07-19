import { Component, OnInit } from '@angular/core';
import { UsuarioService } from '../../services/usuario.service';
import { Usuario } from '../../models/usuario';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-usuario',
  templateUrl: './usuario-view.component.html',
  styleUrls: []
})
export class UsuarioComponent implements OnInit {

  usuarios: Usuario[];
  usuario: Usuario;

  constructor(private usuarioService: UsuarioService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.usuarioService.getUsuarios().then(usuarios => this.usuarios = usuarios);
  }

  newUsuario(): void {

    this.router.navigate(['/usuario/new']).then(() => null);
    this.globals.currentModule = 'Usuario';
  }

  editar(usuario: Usuario): void {
    this.usuario = usuario;
    this.router.navigate(['/usuario/edit', this.usuario._id ]);
  }

  borrar(usuario: Usuario): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar usuario?',
      accept: () => {
        this.usuarioService.delete(usuario._id)
          .then(response => this.usuarioService.getUsuarios().then(usuarios => this.usuarios = usuarios));
      }
    });
  }
}