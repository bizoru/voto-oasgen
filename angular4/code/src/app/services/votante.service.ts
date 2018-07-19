import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Votante } from '../models/votante';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class VotanteService {

  private serviceURL = 'http://localhost:8081/v1/votante';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getVotantes(): Promise<Votante[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Votante[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getVotante(id: string): Promise<Votante> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Votante)
      .catch(this.handleError);
  }


  update(votante: Votante): Promise<Votante> {
    const url = `${this.serviceURL}/${ votante._id}`;
    return this.http
      .put(url, JSON.stringify(votante), {headers: this.headers})
      .toPromise()
      .then(() => votante)
      .catch(this.handleError);
  }


  create(votante: Votante): Promise<Votante> {
    return this.http
      .post(this.serviceURL, JSON.stringify(votante), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Votante)
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