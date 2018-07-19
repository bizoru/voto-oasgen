package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Sufragante
type SufraganteController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Sufragante models.Sufragante
// @Failure 403 :objectId is empty
// @router / [get]
func (j *SufraganteController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllSufragantes(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Sufragante by nombre
// @Param	nombre		path 	string	true		"El nombre de la Sufragante a consultar"
// @Success 200 {object} models.Sufragante
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *SufraganteController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		sufragante, err := models.GetSufraganteById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = sufragante
		}
	}
	j.ServeJSON()
}

// @Title Borrar Sufragante
// @Description Borrar Sufragante
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *SufraganteController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteSufraganteById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Sufragante
// @Description Crear Sufragante
// @Param	body		body 	models.Sufragante	true		"Body para la creacion de Sufragante"
// @Success 200 {int} Sufragante.Id
// @Failure 403 body is empty
// @router / [post]
func (j *SufraganteController) Post() {
	var sufragante models.Sufragante
	json.Unmarshal(j.Ctx.Input.RequestBody, &sufragante)
	fmt.Println(sufragante)
	session,_ := db.GetSession()
	models.InsertSufragante(session,sufragante)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Sufragante
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *SufraganteController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var sufragante models.Sufragante
	json.Unmarshal(j.Ctx.Input.RequestBody, &sufragante)
	session,_ := db.GetSession()

	err := models.UpdateSufragante(session, sufragante,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Sufragante
// @Param	body		body 	models.Sufragante	true		"Body para la creacion de Sufragante"
// @Success 200 {int} Sufragante.Id
// @Failure 403 body is empty
// @router / [options]
func (j *SufraganteController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Sufragante
// @Param	body		body 	models.Sufragante true		"Body para la creacion de Sufragante"
// @Success 200 {int} Sufragante.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *SufraganteController) SufraganteDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}