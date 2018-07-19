package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Configuracion
type ConfiguracionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Configuracion models.Configuracion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ConfiguracionController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllConfiguracions(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Configuracion by nombre
// @Param	nombre		path 	string	true		"El nombre de la Configuracion a consultar"
// @Success 200 {object} models.Configuracion
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ConfiguracionController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		configuracion, err := models.GetConfiguracionById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = configuracion
		}
	}
	j.ServeJSON()
}

// @Title Borrar Configuracion
// @Description Borrar Configuracion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ConfiguracionController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteConfiguracionById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Configuracion
// @Description Crear Configuracion
// @Param	body		body 	models.Configuracion	true		"Body para la creacion de Configuracion"
// @Success 200 {int} Configuracion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ConfiguracionController) Post() {
	var configuracion models.Configuracion
	json.Unmarshal(j.Ctx.Input.RequestBody, &configuracion)
	fmt.Println(configuracion)
	session,_ := db.GetSession()
	models.InsertConfiguracion(session,configuracion)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Configuracion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ConfiguracionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var configuracion models.Configuracion
	json.Unmarshal(j.Ctx.Input.RequestBody, &configuracion)
	session,_ := db.GetSession()

	err := models.UpdateConfiguracion(session, configuracion,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Configuracion
// @Param	body		body 	models.Configuracion	true		"Body para la creacion de Configuracion"
// @Success 200 {int} Configuracion.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ConfiguracionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Configuracion
// @Param	body		body 	models.Configuracion true		"Body para la creacion de Configuracion"
// @Success 200 {int} Configuracion.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ConfiguracionController) ConfiguracionDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}