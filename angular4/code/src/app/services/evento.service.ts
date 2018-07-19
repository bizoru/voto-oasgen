import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Evento } from '../models/evento';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class EventoService {

  private serviceURL = 'http://localhost:8081/v1/evento';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getEventos(): Promise<Evento[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Evento[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getEvento(id: string): Promise<Evento> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Evento)
      .catch(this.handleError);
  }


  update(evento: Evento): Promise<Evento> {
    const url = `${this.serviceURL}/${ evento._id}`;
    return this.http
      .put(url, JSON.stringify(evento), {headers: this.headers})
      .toPromise()
      .then(() => evento)
      .catch(this.handleError);
  }


  create(evento: Evento): Promise<Evento> {
    return this.http
      .post(this.serviceURL, JSON.stringify(evento), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Evento)
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