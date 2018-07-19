import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Historia } from '../models/historia';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class HistoriaService {

  private serviceURL = 'http://localhost:8081/v1/historia';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getHistorias(): Promise<Historia[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Historia[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getHistoria(id: string): Promise<Historia> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Historia)
      .catch(this.handleError);
  }


  update(historia: Historia): Promise<Historia> {
    const url = `${this.serviceURL}/${ historia._id}`;
    return this.http
      .put(url, JSON.stringify(historia), {headers: this.headers})
      .toPromise()
      .then(() => historia)
      .catch(this.handleError);
  }


  create(historia: Historia): Promise<Historia> {
    return this.http
      .post(this.serviceURL, JSON.stringify(historia), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Historia)
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