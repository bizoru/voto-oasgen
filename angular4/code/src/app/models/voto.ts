
import { Candidato } from './candidato';
import { Evento } from './evento';

export class Voto {
   _id: string;
  candidato: Candidato[];
  evento: Evento[];
}