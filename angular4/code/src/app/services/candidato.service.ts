import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Candidato } from '../models/candidato';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class CandidatoService {

  private serviceURL = 'http://localhost:8081/v1/candidato';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getCandidatos(): Promise<Candidato[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Candidato[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getCandidato(id: string): Promise<Candidato> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Candidato)
      .catch(this.handleError);
  }


  update(candidato: Candidato): Promise<Candidato> {
    const url = `${this.serviceURL}/${ candidato._id}`;
    return this.http
      .put(url, JSON.stringify(candidato), {headers: this.headers})
      .toPromise()
      .then(() => candidato)
      .catch(this.handleError);
  }


  create(candidato: Candidato): Promise<Candidato> {
    return this.http
      .post(this.serviceURL, JSON.stringify(candidato), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Candidato)
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