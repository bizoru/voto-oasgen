import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { ProcesoComponent } from './proceso/proceso-view/proceso-view.component';
import { ProcesoNewComponent } from './proceso/proceso-new/proceso-new.component';
import { ProcesoEditComponent } from './proceso/proceso-edit/proceso-edit.component';
import { ProcesoService } from './services/proceso.service';
import { EleccionComponent } from './eleccion/eleccion-view/eleccion-view.component';
import { EleccionNewComponent } from './eleccion/eleccion-new/eleccion-new.component';
import { EleccionEditComponent } from './eleccion/eleccion-edit/eleccion-edit.component';
import { EleccionService } from './services/eleccion.service';
import { TarjetonComponent } from './tarjeton/tarjeton-view/tarjeton-view.component';
import { TarjetonNewComponent } from './tarjeton/tarjeton-new/tarjeton-new.component';
import { TarjetonEditComponent } from './tarjeton/tarjeton-edit/tarjeton-edit.component';
import { TarjetonService } from './services/tarjeton.service';
import { ConfiguracionComponent } from './configuracion/configuracion-view/configuracion-view.component';
import { ConfiguracionNewComponent } from './configuracion/configuracion-new/configuracion-new.component';
import { ConfiguracionEditComponent } from './configuracion/configuracion-edit/configuracion-edit.component';
import { ConfiguracionService } from './services/configuracion.service';
import { CandidatoComponent } from './candidato/candidato-view/candidato-view.component';
import { CandidatoNewComponent } from './candidato/candidato-new/candidato-new.component';
import { CandidatoEditComponent } from './candidato/candidato-edit/candidato-edit.component';
import { CandidatoService } from './services/candidato.service';
import { VotanteComponent } from './votante/votante-view/votante-view.component';
import { VotanteNewComponent } from './votante/votante-new/votante-new.component';
import { VotanteEditComponent } from './votante/votante-edit/votante-edit.component';
import { VotanteService } from './services/votante.service';
import { SufraganteComponent } from './sufragante/sufragante-view/sufragante-view.component';
import { SufraganteNewComponent } from './sufragante/sufragante-new/sufragante-new.component';
import { SufraganteEditComponent } from './sufragante/sufragante-edit/sufragante-edit.component';
import { SufraganteService } from './services/sufragante.service';
import { EventoComponent } from './evento/evento-view/evento-view.component';
import { EventoNewComponent } from './evento/evento-new/evento-new.component';
import { EventoEditComponent } from './evento/evento-edit/evento-edit.component';
import { EventoService } from './services/evento.service';
import { UsuarioComponent } from './usuario/usuario-view/usuario-view.component';
import { UsuarioNewComponent } from './usuario/usuario-new/usuario-new.component';
import { UsuarioEditComponent } from './usuario/usuario-edit/usuario-edit.component';
import { UsuarioService } from './services/usuario.service';
import { HistoriaComponent } from './historia/historia-view/historia-view.component';
import { HistoriaNewComponent } from './historia/historia-new/historia-new.component';
import { HistoriaEditComponent } from './historia/historia-edit/historia-edit.component';
import { HistoriaService } from './services/historia.service';
import { RolComponent } from './rol/rol-view/rol-view.component';
import { RolNewComponent } from './rol/rol-new/rol-new.component';
import { RolEditComponent } from './rol/rol-edit/rol-edit.component';
import { RolService } from './services/rol.service';
import { VariableComponent } from './variable/variable-view/variable-view.component';
import { VariableNewComponent } from './variable/variable-new/variable-new.component';
import { VariableEditComponent } from './variable/variable-edit/variable-edit.component';
import { VariableService } from './services/variable.service';
import { PonderacionComponent } from './ponderacion/ponderacion-view/ponderacion-view.component';
import { PonderacionNewComponent } from './ponderacion/ponderacion-new/ponderacion-new.component';
import { PonderacionEditComponent } from './ponderacion/ponderacion-edit/ponderacion-edit.component';
import { PonderacionService } from './services/ponderacion.service';
import { PartidoComponent } from './partido/partido-view/partido-view.component';
import { PartidoNewComponent } from './partido/partido-new/partido-new.component';
import { PartidoEditComponent } from './partido/partido-edit/partido-edit.component';
import { PartidoService } from './services/partido.service';
import { EstamentoComponent } from './estamento/estamento-view/estamento-view.component';
import { EstamentoNewComponent } from './estamento/estamento-new/estamento-new.component';
import { EstamentoEditComponent } from './estamento/estamento-edit/estamento-edit.component';
import { EstamentoService } from './services/estamento.service';
import { VotoComponent } from './voto/voto-view/voto-view.component';
import { VotoNewComponent } from './voto/voto-new/voto-new.component';
import { VotoEditComponent } from './voto/voto-edit/voto-edit.component';
import { VotoService } from './services/voto.service';
import { CertificadoComponent } from './certificado/certificado-view/certificado-view.component';
import { CertificadoNewComponent } from './certificado/certificado-new/certificado-new.component';
import { CertificadoEditComponent } from './certificado/certificado-edit/certificado-edit.component';
import { CertificadoService } from './services/certificado.service';

import { AppComponent } from './app.component';
import { EncabezadoComponent } from './encabezado/encabezado.component';
import { NavegacionComponent } from './navegacion/navegacion.component';
import { ContenidoComponent } from './contenido/contenido.component';
import { PieComponent } from './pie/pie.component';
import { MigaComponent } from './miga/miga.component';

import { FormsModule } from '@angular/forms'; // Needed to use ngmodel
import { RoutingModule } from './routing/routing.module';
import { HttpModule } from '@angular/http';


// Primeng Modules
import { ButtonModule} from 'primeng/primeng';
import { DataTableModule, SharedModule } from 'primeng/primeng';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { GlobalsComponent } from './globals/globals.component';


import { ConfirmDialogModule, ConfirmationService } from 'primeng/primeng';
import { DialogModule } from 'primeng/primeng';

// Others
import { DatePickerModule } from 'bizoru-datepicker';


@NgModule({
  declarations: [
    AppComponent,
    EncabezadoComponent,
    NavegacionComponent,
    ContenidoComponent,
    PieComponent,
    MigaComponent,
    ProcesoComponent,
    ProcesoNewComponent,
    ProcesoEditComponent,
    EleccionComponent,
    EleccionNewComponent,
    EleccionEditComponent,
    TarjetonComponent,
    TarjetonNewComponent,
    TarjetonEditComponent,
    ConfiguracionComponent,
    ConfiguracionNewComponent,
    ConfiguracionEditComponent,
    CandidatoComponent,
    CandidatoNewComponent,
    CandidatoEditComponent,
    VotanteComponent,
    VotanteNewComponent,
    VotanteEditComponent,
    SufraganteComponent,
    SufraganteNewComponent,
    SufraganteEditComponent,
    EventoComponent,
    EventoNewComponent,
    EventoEditComponent,
    UsuarioComponent,
    UsuarioNewComponent,
    UsuarioEditComponent,
    HistoriaComponent,
    HistoriaNewComponent,
    HistoriaEditComponent,
    RolComponent,
    RolNewComponent,
    RolEditComponent,
    VariableComponent,
    VariableNewComponent,
    VariableEditComponent,
    PonderacionComponent,
    PonderacionNewComponent,
    PonderacionEditComponent,
    PartidoComponent,
    PartidoNewComponent,
    PartidoEditComponent,
    EstamentoComponent,
    EstamentoNewComponent,
    EstamentoEditComponent,
    VotoComponent,
    VotoNewComponent,
    VotoEditComponent,
    CertificadoComponent,
    CertificadoNewComponent,
    CertificadoEditComponent,
    ],
  imports: [
    BrowserModule,
    HttpModule,
    RoutingModule,
    FormsModule,
    ButtonModule,
    DataTableModule,
    SharedModule,
    BrowserAnimationsModule,
    ConfirmDialogModule,
    DialogModule,
    DatePickerModule
  ],
  providers: [
  GlobalsComponent,
  ConfirmationService,
  ProcesoService,
  EleccionService,
  TarjetonService,
  ConfiguracionService,
  CandidatoService,
  VotanteService,
  SufraganteService,
  EventoService,
  UsuarioService,
  HistoriaService,
  RolService,
  VariableService,
  PonderacionService,
  PartidoService,
  EstamentoService,
  VotoService,
  CertificadoService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }