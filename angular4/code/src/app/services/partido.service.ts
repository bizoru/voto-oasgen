import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Partido } from '../models/partido';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class PartidoService {

  private serviceURL = 'http://localhost:8081/v1/partido';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getPartidos(): Promise<Partido[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Partido[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getPartido(id: string): Promise<Partido> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Partido)
      .catch(this.handleError);
  }


  update(partido: Partido): Promise<Partido> {
    const url = `${this.serviceURL}/${ partido._id}`;
    return this.http
      .put(url, JSON.stringify(partido), {headers: this.headers})
      .toPromise()
      .then(() => partido)
      .catch(this.handleError);
  }


  create(partido: Partido): Promise<Partido> {
    return this.http
      .post(this.serviceURL, JSON.stringify(partido), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Partido)
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