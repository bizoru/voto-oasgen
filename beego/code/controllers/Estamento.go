package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Estamento
type EstamentoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Estamento models.Estamento
// @Failure 403 :objectId is empty
// @router / [get]
func (j *EstamentoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllEstamentos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Estamento by nombre
// @Param	nombre		path 	string	true		"El nombre de la Estamento a consultar"
// @Success 200 {object} models.Estamento
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *EstamentoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		estamento, err := models.GetEstamentoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = estamento
		}
	}
	j.ServeJSON()
}

// @Title Borrar Estamento
// @Description Borrar Estamento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *EstamentoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteEstamentoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Estamento
// @Description Crear Estamento
// @Param	body		body 	models.Estamento	true		"Body para la creacion de Estamento"
// @Success 200 {int} Estamento.Id
// @Failure 403 body is empty
// @router / [post]
func (j *EstamentoController) Post() {
	var estamento models.Estamento
	json.Unmarshal(j.Ctx.Input.RequestBody, &estamento)
	fmt.Println(estamento)
	session,_ := db.GetSession()
	models.InsertEstamento(session,estamento)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Estamento
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *EstamentoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var estamento models.Estamento
	json.Unmarshal(j.Ctx.Input.RequestBody, &estamento)
	session,_ := db.GetSession()

	err := models.UpdateEstamento(session, estamento,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Estamento
// @Param	body		body 	models.Estamento	true		"Body para la creacion de Estamento"
// @Success 200 {int} Estamento.Id
// @Failure 403 body is empty
// @router / [options]
func (j *EstamentoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Estamento
// @Param	body		body 	models.Estamento true		"Body para la creacion de Estamento"
// @Success 200 {int} Estamento.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *EstamentoController) EstamentoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}