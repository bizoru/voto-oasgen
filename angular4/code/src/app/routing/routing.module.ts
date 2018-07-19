import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ProcesoComponent } from '../proceso/proceso-view/proceso-view.component';
import { ProcesoNewComponent } from '../proceso/proceso-new/proceso-new.component';
import { ProcesoEditComponent } from '../proceso/proceso-edit/proceso-edit.component';
import { EleccionComponent } from '../eleccion/eleccion-view/eleccion-view.component';
import { EleccionNewComponent } from '../eleccion/eleccion-new/eleccion-new.component';
import { EleccionEditComponent } from '../eleccion/eleccion-edit/eleccion-edit.component';
import { TarjetonComponent } from '../tarjeton/tarjeton-view/tarjeton-view.component';
import { TarjetonNewComponent } from '../tarjeton/tarjeton-new/tarjeton-new.component';
import { TarjetonEditComponent } from '../tarjeton/tarjeton-edit/tarjeton-edit.component';
import { ConfiguracionComponent } from '../configuracion/configuracion-view/configuracion-view.component';
import { ConfiguracionNewComponent } from '../configuracion/configuracion-new/configuracion-new.component';
import { ConfiguracionEditComponent } from '../configuracion/configuracion-edit/configuracion-edit.component';
import { CandidatoComponent } from '../candidato/candidato-view/candidato-view.component';
import { CandidatoNewComponent } from '../candidato/candidato-new/candidato-new.component';
import { CandidatoEditComponent } from '../candidato/candidato-edit/candidato-edit.component';
import { VotanteComponent } from '../votante/votante-view/votante-view.component';
import { VotanteNewComponent } from '../votante/votante-new/votante-new.component';
import { VotanteEditComponent } from '../votante/votante-edit/votante-edit.component';
import { SufraganteComponent } from '../sufragante/sufragante-view/sufragante-view.component';
import { SufraganteNewComponent } from '../sufragante/sufragante-new/sufragante-new.component';
import { SufraganteEditComponent } from '../sufragante/sufragante-edit/sufragante-edit.component';
import { EventoComponent } from '../evento/evento-view/evento-view.component';
import { EventoNewComponent } from '../evento/evento-new/evento-new.component';
import { EventoEditComponent } from '../evento/evento-edit/evento-edit.component';
import { UsuarioComponent } from '../usuario/usuario-view/usuario-view.component';
import { UsuarioNewComponent } from '../usuario/usuario-new/usuario-new.component';
import { UsuarioEditComponent } from '../usuario/usuario-edit/usuario-edit.component';
import { HistoriaComponent } from '../historia/historia-view/historia-view.component';
import { HistoriaNewComponent } from '../historia/historia-new/historia-new.component';
import { HistoriaEditComponent } from '../historia/historia-edit/historia-edit.component';
import { RolComponent } from '../rol/rol-view/rol-view.component';
import { RolNewComponent } from '../rol/rol-new/rol-new.component';
import { RolEditComponent } from '../rol/rol-edit/rol-edit.component';
import { VariableComponent } from '../variable/variable-view/variable-view.component';
import { VariableNewComponent } from '../variable/variable-new/variable-new.component';
import { VariableEditComponent } from '../variable/variable-edit/variable-edit.component';
import { PonderacionComponent } from '../ponderacion/ponderacion-view/ponderacion-view.component';
import { PonderacionNewComponent } from '../ponderacion/ponderacion-new/ponderacion-new.component';
import { PonderacionEditComponent } from '../ponderacion/ponderacion-edit/ponderacion-edit.component';
import { PartidoComponent } from '../partido/partido-view/partido-view.component';
import { PartidoNewComponent } from '../partido/partido-new/partido-new.component';
import { PartidoEditComponent } from '../partido/partido-edit/partido-edit.component';
import { EstamentoComponent } from '../estamento/estamento-view/estamento-view.component';
import { EstamentoNewComponent } from '../estamento/estamento-new/estamento-new.component';
import { EstamentoEditComponent } from '../estamento/estamento-edit/estamento-edit.component';
import { VotoComponent } from '../voto/voto-view/voto-view.component';
import { VotoNewComponent } from '../voto/voto-new/voto-new.component';
import { VotoEditComponent } from '../voto/voto-edit/voto-edit.component';
import { CertificadoComponent } from '../certificado/certificado-view/certificado-view.component';
import { CertificadoNewComponent } from '../certificado/certificado-new/certificado-new.component';
import { CertificadoEditComponent } from '../certificado/certificado-edit/certificado-edit.component';

const routes: Routes = [
  { path: 'proceso', component: ProcesoComponent },
  { path: 'proceso/new', component: ProcesoNewComponent },
  { path: 'proceso/edit/:id', component: ProcesoEditComponent },
  { path: 'eleccion', component: EleccionComponent },
  { path: 'eleccion/new', component: EleccionNewComponent },
  { path: 'eleccion/edit/:id', component: EleccionEditComponent },
  { path: 'tarjeton', component: TarjetonComponent },
  { path: 'tarjeton/new', component: TarjetonNewComponent },
  { path: 'tarjeton/edit/:id', component: TarjetonEditComponent },
  { path: 'configuracion', component: ConfiguracionComponent },
  { path: 'configuracion/new', component: ConfiguracionNewComponent },
  { path: 'configuracion/edit/:id', component: ConfiguracionEditComponent },
  { path: 'candidato', component: CandidatoComponent },
  { path: 'candidato/new', component: CandidatoNewComponent },
  { path: 'candidato/edit/:id', component: CandidatoEditComponent },
  { path: 'votante', component: VotanteComponent },
  { path: 'votante/new', component: VotanteNewComponent },
  { path: 'votante/edit/:id', component: VotanteEditComponent },
  { path: 'sufragante', component: SufraganteComponent },
  { path: 'sufragante/new', component: SufraganteNewComponent },
  { path: 'sufragante/edit/:id', component: SufraganteEditComponent },
  { path: 'evento', component: EventoComponent },
  { path: 'evento/new', component: EventoNewComponent },
  { path: 'evento/edit/:id', component: EventoEditComponent },
  { path: 'usuario', component: UsuarioComponent },
  { path: 'usuario/new', component: UsuarioNewComponent },
  { path: 'usuario/edit/:id', component: UsuarioEditComponent },
  { path: 'historia', component: HistoriaComponent },
  { path: 'historia/new', component: HistoriaNewComponent },
  { path: 'historia/edit/:id', component: HistoriaEditComponent },
  { path: 'rol', component: RolComponent },
  { path: 'rol/new', component: RolNewComponent },
  { path: 'rol/edit/:id', component: RolEditComponent },
  { path: 'variable', component: VariableComponent },
  { path: 'variable/new', component: VariableNewComponent },
  { path: 'variable/edit/:id', component: VariableEditComponent },
  { path: 'ponderacion', component: PonderacionComponent },
  { path: 'ponderacion/new', component: PonderacionNewComponent },
  { path: 'ponderacion/edit/:id', component: PonderacionEditComponent },
  { path: 'partido', component: PartidoComponent },
  { path: 'partido/new', component: PartidoNewComponent },
  { path: 'partido/edit/:id', component: PartidoEditComponent },
  { path: 'estamento', component: EstamentoComponent },
  { path: 'estamento/new', component: EstamentoNewComponent },
  { path: 'estamento/edit/:id', component: EstamentoEditComponent },
  { path: 'voto', component: VotoComponent },
  { path: 'voto/new', component: VotoNewComponent },
  { path: 'voto/edit/:id', component: VotoEditComponent },
  { path: 'certificado', component: CertificadoComponent },
  { path: 'certificado/new', component: CertificadoNewComponent },
  { path: 'certificado/edit/:id', component: CertificadoEditComponent },
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forRoot(routes)
  ],
  exports: [RouterModule],
  declarations: []
})
export class RoutingModule { }