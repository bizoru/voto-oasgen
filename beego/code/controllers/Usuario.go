package controllers

import (
  "api/db"
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_ "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Usuario
type UsuarioController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 Usuario models.Usuario
// @Failure 403 :objectId is empty
// @router / [get]
func (j *UsuarioController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllUsuarios(session)

  if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// @Title Get
// @Description get Usuario by nombre
// @Param	nombre		path 	string	true		"El nombre de la Usuario a consultar"
// @Success 200 {object} models.Usuario
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *UsuarioController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		usuario, err := models.GetUsuarioById(session,id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = usuario
		}
	}
	j.ServeJSON()
}

// @Title Borrar Usuario
// @Description Borrar Usuario
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *UsuarioController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteUsuarioById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Usuario
// @Description Crear Usuario
// @Param	body		body 	models.Usuario	true		"Body para la creacion de Usuario"
// @Success 200 {int} Usuario.Id
// @Failure 403 body is empty
// @router / [post]
func (j *UsuarioController) Post() {
	var usuario models.Usuario
	json.Unmarshal(j.Ctx.Input.RequestBody, &usuario)
	fmt.Println(usuario)
	session,_ := db.GetSession()
	models.InsertUsuario(session,usuario)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Usuario
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *UsuarioController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var usuario models.Usuario
	json.Unmarshal(j.Ctx.Input.RequestBody, &usuario)
	session,_ := db.GetSession()

	err := models.UpdateUsuario(session, usuario,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Usuario
// @Param	body		body 	models.Usuario	true		"Body para la creacion de Usuario"
// @Success 200 {int} Usuario.Id
// @Failure 403 body is empty
// @router / [options]
func (j *UsuarioController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear Usuario
// @Param	body		body 	models.Usuario true		"Body para la creacion de Usuario"
// @Success 200 {int} Usuario.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *UsuarioController) UsuarioDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}