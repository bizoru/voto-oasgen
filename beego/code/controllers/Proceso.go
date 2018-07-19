package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Proceso
type ProcesoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Proceso models.Proceso
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ProcesoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllProcesos(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Proceso by nombre
// @Param	nombre		path 	string	true		"El nombre de la Proceso a consultar"
// @Success 200 {object} models.Proceso
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ProcesoController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		proceso, err := models.GetProcesoById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = proceso
		}
	}
	j.ServeJSON()
}

// @Title Borrar Proceso
// @Description Borrar Proceso
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ProcesoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteProcesoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Proceso
// @Description Crear Proceso
// @Param	body		body 	models.Proceso	true		"Body para la creacion de Proceso"
// @Success 200 {int} Proceso.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ProcesoController) Post() {
	var proceso models.Proceso
	json.Unmarshal(j.Ctx.Input.RequestBody, &proceso)
	fmt.Println(proceso)
	session,_ := db.GetSession()
	models.InsertProceso(session,proceso)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Proceso
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ProcesoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var proceso models.Proceso
	json.Unmarshal(j.Ctx.Input.RequestBody, &proceso)
	session,_ := db.GetSession()

	err := models.UpdateProceso(session, proceso,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Proceso
// @Param	body		body 	models.Proceso	true		"Body para la creacion de Proceso"
// @Success 200 {int} Proceso.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ProcesoController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Proceso
// @Param	body		body 	models.Proceso true		"Body para la creacion de Proceso"
// @Success 200 {int} Proceso.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ProcesoController) ProcesoDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}