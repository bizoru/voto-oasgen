import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Estamento } from '../models/estamento';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class EstamentoService {

  private serviceURL = 'http://localhost:8081/v1/estamento';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getEstamentos(): Promise<Estamento[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Estamento[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getEstamento(id: string): Promise<Estamento> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Estamento)
      .catch(this.handleError);
  }


  update(estamento: Estamento): Promise<Estamento> {
    const url = `${this.serviceURL}/${ estamento._id}`;
    return this.http
      .put(url, JSON.stringify(estamento), {headers: this.headers})
      .toPromise()
      .then(() => estamento)
      .catch(this.handleError);
  }


  create(estamento: Estamento): Promise<Estamento> {
    return this.http
      .post(this.serviceURL, JSON.stringify(estamento), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Estamento)
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