import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Ponderacion } from '../models/ponderacion';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class PonderacionService {

  private serviceURL = 'http://localhost:8081/v1/ponderacion';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getPonderacions(): Promise<Ponderacion[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Ponderacion[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getPonderacion(id: string): Promise<Ponderacion> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Ponderacion)
      .catch(this.handleError);
  }


  update(ponderacion: Ponderacion): Promise<Ponderacion> {
    const url = `${this.serviceURL}/${ ponderacion._id}`;
    return this.http
      .put(url, JSON.stringify(ponderacion), {headers: this.headers})
      .toPromise()
      .then(() => ponderacion)
      .catch(this.handleError);
  }


  create(ponderacion: Ponderacion): Promise<Ponderacion> {
    return this.http
      .post(this.serviceURL, JSON.stringify(ponderacion), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Ponderacion)
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