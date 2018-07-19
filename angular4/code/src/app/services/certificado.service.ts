import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Certificado } from '../models/certificado';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class CertificadoService {

  private serviceURL = 'http://localhost:8081/v1/certificado';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getCertificados(): Promise<Certificado[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Certificado[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getCertificado(id: string): Promise<Certificado> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Certificado)
      .catch(this.handleError);
  }


  update(certificado: Certificado): Promise<Certificado> {
    const url = `${this.serviceURL}/${ certificado._id}`;
    return this.http
      .put(url, JSON.stringify(certificado), {headers: this.headers})
      .toPromise()
      .then(() => certificado)
      .catch(this.handleError);
  }


  create(certificado: Certificado): Promise<Certificado> {
    return this.http
      .post(this.serviceURL, JSON.stringify(certificado), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Certificado)
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