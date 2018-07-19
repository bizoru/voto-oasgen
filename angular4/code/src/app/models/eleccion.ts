
import { Votante } from './votante';
import { Ponderacion } from './ponderacion';
import { Usuario } from './usuario';

export class Eleccion {
   _id: string;
  nombre:	string;
  votante: Votante[];
  ponderacion: Ponderacion[];
  responsable: Usuario[];
  habilitado:	boolean;
}