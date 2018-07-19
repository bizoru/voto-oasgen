import { Component, OnInit } from '@angular/core';
import { Voto } from '../../models/voto';
import { VotoService } from '../../services/voto.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-voto-edit',
  templateUrl: './voto-edit.component.html',
  styleUrls: []
})
export class VotoEditComponent implements OnInit {

  voto: Voto = new Voto();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private votoService: VotoService) {

  }

  actualizar(voto: Voto): void {
    this.votoService.update(voto).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.votoService.getVoto(params['id']))
      .subscribe(voto => this.voto = voto);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}