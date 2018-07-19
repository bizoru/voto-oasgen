
import { Proceso } from './proceso';
import { Eleccion } from './eleccion';
import { Sufragante } from './sufragante';

export class Certificado {
   _id: string;
  proceso: Proceso[];
  eleccion: Eleccion[];
  sufragante: Sufragante[];
}