package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Certificado
type CertificadoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Certificado models.Certificado
// @Failure 403 :objectId is empty
// @router / [get]
func (j *CertificadoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllCertificados(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Certificado by nombre
// @Param	nombre		path 	string	true		"El nombre de la Certificado a consultar"
// @Success 200 {object} models.Certificado
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *CertificadoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		certificado, err := models.GetCertificadoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = certificado
		}
	}
	j.ServeJSON()
}

// @Title Borrar Certificado
// @Description Borrar Certificado
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *CertificadoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteCertificadoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Certificado
// @Description Crear Certificado
// @Param	body		body 	models.Certificado	true		"Body para la creacion de Certificado"
// @Success 200 {int} Certificado.Id
// @Failure 403 body is empty
// @router / [post]
func (j *CertificadoController) Post() {
	var certificado models.Certificado
	json.Unmarshal(j.Ctx.Input.RequestBody, &certificado)
	fmt.Println(certificado)
	session,_ := db.GetSession()
	models.InsertCertificado(session,certificado)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Certificado
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *CertificadoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var certificado models.Certificado
	json.Unmarshal(j.Ctx.Input.RequestBody, &certificado)
	session,_ := db.GetSession()

	err := models.UpdateCertificado(session, certificado,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Certificado
// @Param	body		body 	models.Certificado	true		"Body para la creacion de Certificado"
// @Success 200 {int} Certificado.Id
// @Failure 403 body is empty
// @router / [options]
func (j *CertificadoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Certificado
// @Param	body		body 	models.Certificado true		"Body para la creacion de Certificado"
// @Success 200 {int} Certificado.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *CertificadoController) CertificadoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}