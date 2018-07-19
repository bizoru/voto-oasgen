import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Usuario } from '../models/usuario';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class UsuarioService {

  private serviceURL = 'http://localhost:8081/v1/usuario';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getUsuarios(): Promise<Usuario[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Usuario[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getUsuario(id: string): Promise<Usuario> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Usuario)
      .catch(this.handleError);
  }


  update(usuario: Usuario): Promise<Usuario> {
    const url = `${this.serviceURL}/${ usuario._id}`;
    return this.http
      .put(url, JSON.stringify(usuario), {headers: this.headers})
      .toPromise()
      .then(() => usuario)
      .catch(this.handleError);
  }


  create(usuario: Usuario): Promise<Usuario> {
    return this.http
      .post(this.serviceURL, JSON.stringify(usuario), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Usuario)
      .catch(this.handleError);
  }

  delete(id: string): Promise<void> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.delete(url, {headers: this.headers})
      .toPromise()
      .then(() => null)
      .catch(this.handleError);
  }

}