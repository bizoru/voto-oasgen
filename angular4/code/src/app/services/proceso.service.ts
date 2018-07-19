import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Proceso } from '../models/proceso';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class ProcesoService {

  private serviceURL = 'http://localhost:8081/v1/proceso';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getProcesos(): Promise<Proceso[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Proceso[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getProceso(id: string): Promise<Proceso> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Proceso)
      .catch(this.handleError);
  }


  update(proceso: Proceso): Promise<Proceso> {
    const url = `${this.serviceURL}/${ proceso._id}`;
    return this.http
      .put(url, JSON.stringify(proceso), {headers: this.headers})
      .toPromise()
      .then(() => proceso)
      .catch(this.handleError);
  }


  create(proceso: Proceso): Promise<Proceso> {
    return this.http
      .post(this.serviceURL, JSON.stringify(proceso), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Proceso)
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