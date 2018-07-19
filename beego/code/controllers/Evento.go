package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Evento
type EventoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Evento models.Evento
// @Failure 403 :objectId is empty
// @router / [get]
func (j *EventoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllEventos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Evento by nombre
// @Param	nombre		path 	string	true		"El nombre de la Evento a consultar"
// @Success 200 {object} models.Evento
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *EventoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		evento, err := models.GetEventoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = evento
		}
	}
	j.ServeJSON()
}

// @Title Borrar Evento
// @Description Borrar Evento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *EventoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteEventoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Evento
// @Description Crear Evento
// @Param	body		body 	models.Evento	true		"Body para la creacion de Evento"
// @Success 200 {int} Evento.Id
// @Failure 403 body is empty
// @router / [post]
func (j *EventoController) Post() {
	var evento models.Evento
	json.Unmarshal(j.Ctx.Input.RequestBody, &evento)
	fmt.Println(evento)
	session,_ := db.GetSession()
	models.InsertEvento(session,evento)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Evento
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *EventoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var evento models.Evento
	json.Unmarshal(j.Ctx.Input.RequestBody, &evento)
	session,_ := db.GetSession()

	err := models.UpdateEvento(session, evento,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Evento
// @Param	body		body 	models.Evento	true		"Body para la creacion de Evento"
// @Success 200 {int} Evento.Id
// @Failure 403 body is empty
// @router / [options]
func (j *EventoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Evento
// @Param	body		body 	models.Evento true		"Body para la creacion de Evento"
// @Success 200 {int} Evento.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *EventoController) EventoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}