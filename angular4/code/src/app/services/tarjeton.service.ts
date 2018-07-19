import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Tarjeton } from '../models/tarjeton';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class TarjetonService {

  private serviceURL = 'http://localhost:8081/v1/tarjeton';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getTarjetons(): Promise<Tarjeton[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Tarjeton[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getTarjeton(id: string): Promise<Tarjeton> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Tarjeton)
      .catch(this.handleError);
  }


  update(tarjeton: Tarjeton): Promise<Tarjeton> {
    const url = `${this.serviceURL}/${ tarjeton._id}`;
    return this.http
      .put(url, JSON.stringify(tarjeton), {headers: this.headers})
      .toPromise()
      .then(() => tarjeton)
      .catch(this.handleError);
  }


  create(tarjeton: Tarjeton): Promise<Tarjeton> {
    return this.http
      .post(this.serviceURL, JSON.stringify(tarjeton), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Tarjeton)
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