import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Sufragante } from '../models/sufragante';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class SufraganteService {

  private serviceURL = 'http://localhost:8081/v1/sufragante';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getSufragantes(): Promise<Sufragante[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Sufragante[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getSufragante(id: string): Promise<Sufragante> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Sufragante)
      .catch(this.handleError);
  }


  update(sufragante: Sufragante): Promise<Sufragante> {
    const url = `${this.serviceURL}/${ sufragante._id}`;
    return this.http
      .put(url, JSON.stringify(sufragante), {headers: this.headers})
      .toPromise()
      .then(() => sufragante)
      .catch(this.handleError);
  }


  create(sufragante: Sufragante): Promise<Sufragante> {
    return this.http
      .post(this.serviceURL, JSON.stringify(sufragante), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Sufragante)
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