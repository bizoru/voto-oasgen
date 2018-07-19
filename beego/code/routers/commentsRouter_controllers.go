package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CandidatoController"] = append(beego.GlobalControllerRouter["api/controllers:CandidatoController"],
		beego.ControllerComments{
			Method: "CandidatoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["api/controllers:CertificadoController"],
		beego.ControllerComments{
			Method: "CertificadoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ConfiguracionController"] = append(beego.GlobalControllerRouter["api/controllers:ConfiguracionController"],
		beego.ControllerComments{
			Method: "ConfiguracionDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EleccionController"] = append(beego.GlobalControllerRouter["api/controllers:EleccionController"],
		beego.ControllerComments{
			Method: "EleccionDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EstamentoController"] = append(beego.GlobalControllerRouter["api/controllers:EstamentoController"],
		beego.ControllerComments{
			Method: "EstamentoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:EventoController"] = append(beego.GlobalControllerRouter["api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "EventoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:HistoriaController"] = append(beego.GlobalControllerRouter["api/controllers:HistoriaController"],
		beego.ControllerComments{
			Method: "HistoriaDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PartidoController"] = append(beego.GlobalControllerRouter["api/controllers:PartidoController"],
		beego.ControllerComments{
			Method: "PartidoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:PonderacionController"] = append(beego.GlobalControllerRouter["api/controllers:PonderacionController"],
		beego.ControllerComments{
			Method: "PonderacionDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["api/controllers:ProcesoController"],
		beego.ControllerComments{
			Method: "ProcesoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:RolController"] = append(beego.GlobalControllerRouter["api/controllers:RolController"],
		beego.ControllerComments{
			Method: "RolDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:SufraganteController"] = append(beego.GlobalControllerRouter["api/controllers:SufraganteController"],
		beego.ControllerComments{
			Method: "SufraganteDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:TarjetonController"] = append(beego.GlobalControllerRouter["api/controllers:TarjetonController"],
		beego.ControllerComments{
			Method: "TarjetonDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method: "UsuarioDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VariableController"] = append(beego.GlobalControllerRouter["api/controllers:VariableController"],
		beego.ControllerComments{
			Method: "VariableDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotanteController"] = append(beego.GlobalControllerRouter["api/controllers:VotanteController"],
		beego.ControllerComments{
			Method: "VotanteDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:VotoController"] = append(beego.GlobalControllerRouter["api/controllers:VotoController"],
		beego.ControllerComments{
			Method: "VotoDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

}
