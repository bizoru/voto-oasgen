// @APIVersion 1.0.0
// @Title API
// @Description API Aplicacion Voto - Entidades Core
// @Contact ssierraf@correo.udistrital.edu.co
// @TermsOfServiceUrl http://oas.udistrital.edu.co/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
  beego.NSNamespace("/proceso",
	    beego.NSInclude(
	      &controllers.ProcesoController{},
	    ),
   ),
  beego.NSNamespace("/eleccion",
	    beego.NSInclude(
	      &controllers.EleccionController{},
	    ),
   ),
  beego.NSNamespace("/tarjeton",
	    beego.NSInclude(
	      &controllers.TarjetonController{},
	    ),
   ),
  beego.NSNamespace("/configuracion",
	    beego.NSInclude(
	      &controllers.ConfiguracionController{},
	    ),
   ),
  beego.NSNamespace("/candidato",
	    beego.NSInclude(
	      &controllers.CandidatoController{},
	    ),
   ),
  beego.NSNamespace("/votante",
	    beego.NSInclude(
	      &controllers.VotanteController{},
	    ),
   ),
  beego.NSNamespace("/sufragante",
	    beego.NSInclude(
	      &controllers.SufraganteController{},
	    ),
   ),
  beego.NSNamespace("/evento",
	    beego.NSInclude(
	      &controllers.EventoController{},
	    ),
   ),
  beego.NSNamespace("/usuario",
	    beego.NSInclude(
	      &controllers.UsuarioController{},
	    ),
   ),
  beego.NSNamespace("/historia",
	    beego.NSInclude(
	      &controllers.HistoriaController{},
	    ),
   ),
  beego.NSNamespace("/rol",
	    beego.NSInclude(
	      &controllers.RolController{},
	    ),
   ),
  beego.NSNamespace("/variable",
	    beego.NSInclude(
	      &controllers.VariableController{},
	    ),
   ),
  beego.NSNamespace("/ponderacion",
	    beego.NSInclude(
	      &controllers.PonderacionController{},
	    ),
   ),
  beego.NSNamespace("/partido",
	    beego.NSInclude(
	      &controllers.PartidoController{},
	    ),
   ),
  beego.NSNamespace("/estamento",
	    beego.NSInclude(
	      &controllers.EstamentoController{},
	    ),
   ),
  beego.NSNamespace("/voto",
	    beego.NSInclude(
	      &controllers.VotoController{},
	    ),
   ),
  beego.NSNamespace("/certificado",
	    beego.NSInclude(
	      &controllers.CertificadoController{},
	    ),
   ),
	)
	beego.AddNamespace(ns)
}