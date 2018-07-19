import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Voto } from '../models/voto';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class VotoService {

  private serviceURL = 'http://localhost:8081/v1/voto';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getVotos(): Promise<Voto[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Voto[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getVoto(id: string): Promise<Voto> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Voto)
      .catch(this.handleError);
  }


  update(voto: Voto): Promise<Voto> {
    const url = `${this.serviceURL}/${ voto._id}`;
    return this.http
      .put(url, JSON.stringify(voto), {headers: this.headers})
      .toPromise()
      .then(() => voto)
      .catch(this.handleError);
  }


  create(voto: Voto): Promise<Voto> {
    return this.http
      .post(this.serviceURL, JSON.stringify(voto), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Voto)
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