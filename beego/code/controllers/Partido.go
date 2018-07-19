package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Partido
type PartidoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Partido models.Partido
// @Failure 403 :objectId is empty
// @router / [get]
func (j *PartidoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllPartidos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Partido by nombre
// @Param	nombre		path 	string	true		"El nombre de la Partido a consultar"
// @Success 200 {object} models.Partido
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *PartidoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		partido, err := models.GetPartidoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = partido
		}
	}
	j.ServeJSON()
}

// @Title Borrar Partido
// @Description Borrar Partido
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *PartidoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeletePartidoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Partido
// @Description Crear Partido
// @Param	body		body 	models.Partido	true		"Body para la creacion de Partido"
// @Success 200 {int} Partido.Id
// @Failure 403 body is empty
// @router / [post]
func (j *PartidoController) Post() {
	var partido models.Partido
	json.Unmarshal(j.Ctx.Input.RequestBody, &partido)
	fmt.Println(partido)
	session,_ := db.GetSession()
	models.InsertPartido(session,partido)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Partido
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *PartidoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var partido models.Partido
	json.Unmarshal(j.Ctx.Input.RequestBody, &partido)
	session,_ := db.GetSession()

	err := models.UpdatePartido(session, partido,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Partido
// @Param	body		body 	models.Partido	true		"Body para la creacion de Partido"
// @Success 200 {int} Partido.Id
// @Failure 403 body is empty
// @router / [options]
func (j *PartidoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Partido
// @Param	body		body 	models.Partido true		"Body para la creacion de Partido"
// @Success 200 {int} Partido.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *PartidoController) PartidoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}