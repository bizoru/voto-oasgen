import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Configuracion } from '../models/configuracion';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class ConfiguracionService {

  private serviceURL = 'http://localhost:8081/v1/configuracion';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getConfiguracions(): Promise<Configuracion[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Configuracion[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getConfiguracion(id: string): Promise<Configuracion> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Configuracion)
      .catch(this.handleError);
  }


  update(configuracion: Configuracion): Promise<Configuracion> {
    const url = `${this.serviceURL}/${ configuracion._id}`;
    return this.http
      .put(url, JSON.stringify(configuracion), {headers: this.headers})
      .toPromise()
      .then(() => configuracion)
      .catch(this.handleError);
  }


  create(configuracion: Configuracion): Promise<Configuracion> {
    return this.http
      .post(this.serviceURL, JSON.stringify(configuracion), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Configuracion)
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