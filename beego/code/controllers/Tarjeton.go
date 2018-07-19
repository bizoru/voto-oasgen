package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Tarjeton
type TarjetonController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Tarjeton models.Tarjeton
// @Failure 403 :objectId is empty
// @router / [get]
func (j *TarjetonController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllTarjetons(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Tarjeton by nombre
// @Param	nombre		path 	string	true		"El nombre de la Tarjeton a consultar"
// @Success 200 {object} models.Tarjeton
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *TarjetonController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		tarjeton, err := models.GetTarjetonById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = tarjeton
		}
	}
	j.ServeJSON()
}

// @Title Borrar Tarjeton
// @Description Borrar Tarjeton
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *TarjetonController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteTarjetonById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Tarjeton
// @Description Crear Tarjeton
// @Param	body		body 	models.Tarjeton	true		"Body para la creacion de Tarjeton"
// @Success 200 {int} Tarjeton.Id
// @Failure 403 body is empty
// @router / [post]
func (j *TarjetonController) Post() {
	var tarjeton models.Tarjeton
	json.Unmarshal(j.Ctx.Input.RequestBody, &tarjeton)
	fmt.Println(tarjeton)
	session,_ := db.GetSession()
	models.InsertTarjeton(session,tarjeton)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Tarjeton
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *TarjetonController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var tarjeton models.Tarjeton
	json.Unmarshal(j.Ctx.Input.RequestBody, &tarjeton)
	session,_ := db.GetSession()

	err := models.UpdateTarjeton(session, tarjeton,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Tarjeton
// @Param	body		body 	models.Tarjeton	true		"Body para la creacion de Tarjeton"
// @Success 200 {int} Tarjeton.Id
// @Failure 403 body is empty
// @router / [options]
func (j *TarjetonController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Tarjeton
// @Param	body		body 	models.Tarjeton true		"Body para la creacion de Tarjeton"
// @Success 200 {int} Tarjeton.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *TarjetonController) TarjetonDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}