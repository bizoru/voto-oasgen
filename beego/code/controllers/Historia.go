package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Historia
type HistoriaController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Historia models.Historia
// @Failure 403 :objectId is empty
// @router / [get]
func (j *HistoriaController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllHistorias(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Historia by nombre
// @Param	nombre		path 	string	true		"El nombre de la Historia a consultar"
// @Success 200 {object} models.Historia
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *HistoriaController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		historia, err := models.GetHistoriaById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = historia
		}
	}
	j.ServeJSON()
}

// @Title Borrar Historia
// @Description Borrar Historia
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *HistoriaController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteHistoriaById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Historia
// @Description Crear Historia
// @Param	body		body 	models.Historia	true		"Body para la creacion de Historia"
// @Success 200 {int} Historia.Id
// @Failure 403 body is empty
// @router / [post]
func (j *HistoriaController) Post() {
	var historia models.Historia
	json.Unmarshal(j.Ctx.Input.RequestBody, &historia)
	fmt.Println(historia)
	session,_ := db.GetSession()
	models.InsertHistoria(session,historia)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Historia
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *HistoriaController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var historia models.Historia
	json.Unmarshal(j.Ctx.Input.RequestBody, &historia)
	session,_ := db.GetSession()

	err := models.UpdateHistoria(session, historia,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Historia
// @Param	body		body 	models.Historia	true		"Body para la creacion de Historia"
// @Success 200 {int} Historia.Id
// @Failure 403 body is empty
// @router / [options]
func (j *HistoriaController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Historia
// @Param	body		body 	models.Historia true		"Body para la creacion de Historia"
// @Success 200 {int} Historia.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *HistoriaController) HistoriaDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}