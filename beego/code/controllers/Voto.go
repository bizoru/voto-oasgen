package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Voto
type VotoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Voto models.Voto
// @Failure 403 :objectId is empty
// @router / [get]
func (j *VotoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllVotos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Voto by nombre
// @Param	nombre		path 	string	true		"El nombre de la Voto a consultar"
// @Success 200 {object} models.Voto
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *VotoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		voto, err := models.GetVotoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = voto
		}
	}
	j.ServeJSON()
}

// @Title Borrar Voto
// @Description Borrar Voto
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *VotoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteVotoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Voto
// @Description Crear Voto
// @Param	body		body 	models.Voto	true		"Body para la creacion de Voto"
// @Success 200 {int} Voto.Id
// @Failure 403 body is empty
// @router / [post]
func (j *VotoController) Post() {
	var voto models.Voto
	json.Unmarshal(j.Ctx.Input.RequestBody, &voto)
	fmt.Println(voto)
	session,_ := db.GetSession()
	models.InsertVoto(session,voto)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Voto
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *VotoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var voto models.Voto
	json.Unmarshal(j.Ctx.Input.RequestBody, &voto)
	session,_ := db.GetSession()

	err := models.UpdateVoto(session, voto,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Voto
// @Param	body		body 	models.Voto	true		"Body para la creacion de Voto"
// @Success 200 {int} Voto.Id
// @Failure 403 body is empty
// @router / [options]
func (j *VotoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Voto
// @Param	body		body 	models.Voto true		"Body para la creacion de Voto"
// @Success 200 {int} Voto.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *VotoController) VotoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}