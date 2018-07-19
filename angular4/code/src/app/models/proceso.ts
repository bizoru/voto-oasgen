
import { Eleccion } from './eleccion';

export class Proceso {
   _id: string;
  eleccion: Eleccion[];
  nombre:	string;
  fecha_inicio:	Date;
  fecha_final:	Date;
}