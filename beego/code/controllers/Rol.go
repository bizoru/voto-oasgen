package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Rol
type RolController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Rol models.Rol
// @Failure 403 :objectId is empty
// @router / [get]
func (j *RolController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllRols(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Rol by nombre
// @Param	nombre		path 	string	true		"El nombre de la Rol a consultar"
// @Success 200 {object} models.Rol
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *RolController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		rol, err := models.GetRolById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = rol
		}
	}
	j.ServeJSON()
}

// @Title Borrar Rol
// @Description Borrar Rol
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *RolController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteRolById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Rol
// @Description Crear Rol
// @Param	body		body 	models.Rol	true		"Body para la creacion de Rol"
// @Success 200 {int} Rol.Id
// @Failure 403 body is empty
// @router / [post]
func (j *RolController) Post() {
	var rol models.Rol
	json.Unmarshal(j.Ctx.Input.RequestBody, &rol)
	fmt.Println(rol)
	session,_ := db.GetSession()
	models.InsertRol(session,rol)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Rol
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *RolController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var rol models.Rol
	json.Unmarshal(j.Ctx.Input.RequestBody, &rol)
	session,_ := db.GetSession()

	err := models.UpdateRol(session, rol,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Rol
// @Param	body		body 	models.Rol	true		"Body para la creacion de Rol"
// @Success 200 {int} Rol.Id
// @Failure 403 body is empty
// @router / [options]
func (j *RolController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Rol
// @Param	body		body 	models.Rol true		"Body para la creacion de Rol"
// @Success 200 {int} Rol.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *RolController) RolDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}