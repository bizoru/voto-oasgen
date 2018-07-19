import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Variable } from '../models/variable';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class VariableService {

  private serviceURL = 'http://localhost:8081/v1/variable';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getVariables(): Promise<Variable[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Variable[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getVariable(id: string): Promise<Variable> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Variable)
      .catch(this.handleError);
  }


  update(variable: Variable): Promise<Variable> {
    const url = `${this.serviceURL}/${ variable._id}`;
    return this.http
      .put(url, JSON.stringify(variable), {headers: this.headers})
      .toPromise()
      .then(() => variable)
      .catch(this.handleError);
  }


  create(variable: Variable): Promise<Variable> {
    return this.http
      .post(this.serviceURL, JSON.stringify(variable), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Variable)
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