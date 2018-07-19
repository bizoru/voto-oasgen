package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Candidato
type CandidatoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Candidato models.Candidato
// @Failure 403 :objectId is empty
// @router / [get]
func (j *CandidatoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllCandidatos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Candidato by nombre
// @Param	nombre		path 	string	true		"El nombre de la Candidato a consultar"
// @Success 200 {object} models.Candidato
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *CandidatoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		candidato, err := models.GetCandidatoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = candidato
		}
	}
	j.ServeJSON()
}

// @Title Borrar Candidato
// @Description Borrar Candidato
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *CandidatoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteCandidatoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Candidato
// @Description Crear Candidato
// @Param	body		body 	models.Candidato	true		"Body para la creacion de Candidato"
// @Success 200 {int} Candidato.Id
// @Failure 403 body is empty
// @router / [post]
func (j *CandidatoController) Post() {
	var candidato models.Candidato
	json.Unmarshal(j.Ctx.Input.RequestBody, &candidato)
	fmt.Println(candidato)
	session,_ := db.GetSession()
	models.InsertCandidato(session,candidato)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Candidato
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *CandidatoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var candidato models.Candidato
	json.Unmarshal(j.Ctx.Input.RequestBody, &candidato)
	session,_ := db.GetSession()

	err := models.UpdateCandidato(session, candidato,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Candidato
// @Param	body		body 	models.Candidato	true		"Body para la creacion de Candidato"
// @Success 200 {int} Candidato.Id
// @Failure 403 body is empty
// @router / [options]
func (j *CandidatoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Candidato
// @Param	body		body 	models.Candidato true		"Body para la creacion de Candidato"
// @Success 200 {int} Candidato.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *CandidatoController) CandidatoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}