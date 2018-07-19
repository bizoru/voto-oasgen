
import { Estamento } from './estamento';

export class Votante {
   _id: string;
  identificacion:	string;
  nombre:	string;
  apellidos:	string;
  codigo:	string;
  estamento: Estamento[];
}