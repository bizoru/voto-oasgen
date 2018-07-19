import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Eleccion } from '../models/eleccion';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class EleccionService {

  private serviceURL = 'http://localhost:8081/v1/eleccion';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getEleccions(): Promise<Eleccion[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Eleccion[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getEleccion(id: string): Promise<Eleccion> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Eleccion)
      .catch(this.handleError);
  }


  update(eleccion: Eleccion): Promise<Eleccion> {
    const url = `${this.serviceURL}/${ eleccion._id}`;
    return this.http
      .put(url, JSON.stringify(eleccion), {headers: this.headers})
      .toPromise()
      .then(() => eleccion)
      .catch(this.handleError);
  }


  create(eleccion: Eleccion): Promise<Eleccion> {
    return this.http
      .post(this.serviceURL, JSON.stringify(eleccion), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Eleccion)
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