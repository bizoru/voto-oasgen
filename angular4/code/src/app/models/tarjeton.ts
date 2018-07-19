
import { Candidato } from './candidato';
import { Eleccion } from './eleccion';
import { Configuracion } from './configuracion';

export class Tarjeton {
   _id: string;
  candidato: Candidato[];
  eleccion: Eleccion[];
  configuracion: Configuracion[];
}