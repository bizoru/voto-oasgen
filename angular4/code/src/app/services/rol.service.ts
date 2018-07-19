import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Rol } from '../models/rol';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class RolService {

  private serviceURL = 'http://localhost:8081/v1/rol';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getRols(): Promise<Rol[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Rol[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getRol(id: string): Promise<Rol> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Rol)
      .catch(this.handleError);
  }


  update(rol: Rol): Promise<Rol> {
    const url = `${this.serviceURL}/${ rol._id}`;
    return this.http
      .put(url, JSON.stringify(rol), {headers: this.headers})
      .toPromise()
      .then(() => rol)
      .catch(this.handleError);
  }


  create(rol: Rol): Promise<Rol> {
    return this.http
      .post(this.serviceURL, JSON.stringify(rol), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Rol)
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