package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const RolCollection = "rol"

type Rol struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Modulo string `json:"modulo"`
  Permisos string `json:"permisos"`
}


func UpdateRol(session *mgo.Session, j Rol, id string) error{
	c := db.Cursor(session,RolCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertRol(session *mgo.Session, j Rol) {
	c := db.Cursor(session,RolCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllRols(session *mgo.Session) []Rol {
	c := db.Cursor(session,RolCollection)
	defer session.Close()
    fmt.Println("Getting all rols")
	var rols []Rol
	err := c.Find(bson.M{}).All(&rols)
	if err != nil {
		fmt.Println(err)
	}
	return rols
}

func GetRolById(session *mgo.Session,id string) ([]Rol,error) {
	c := db.Cursor(session, RolCollection)
	defer session.Close()
	var rols []Rol
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&rols)
	return rols,err
}

func DeleteRolById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, RolCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}