package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const UsuarioCollection = "usuario"

type Usuario struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
  Apellidos string `json:"apellidos"`
  Identificacion string `json:"identificacion"`
  Roles string `json:"roles"`
  Contrasena string `json:"contrasena"`
}


func UpdateUsuario(session *mgo.Session, j Usuario, id string) error{
	c := db.Cursor(session,UsuarioCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertUsuario(session *mgo.Session, j Usuario) {
	c := db.Cursor(session,UsuarioCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllUsuarios(session *mgo.Session) []Usuario {
	c := db.Cursor(session,UsuarioCollection)
	defer session.Close()
    fmt.Println("Getting all usuarios")
	var usuarios []Usuario
	err := c.Find(bson.M{}).All(&usuarios)
	if err != nil {
		fmt.Println(err)
	}
	return usuarios
}

func GetUsuarioById(session *mgo.Session,id string) ([]Usuario,error) {
	c := db.Cursor(session, UsuarioCollection)
	defer session.Close()
	var usuarios []Usuario
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&usuarios)
	return usuarios,err
}

func DeleteUsuarioById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, UsuarioCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}