package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const ConfiguracionCollection = "configuracion"

type Configuracion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Columnas int `json:"columnas"`
  Filas int `json:"filas"`
}


func UpdateConfiguracion(session *mgo.Session, j Configuracion, id string) error{
	c := db.Cursor(session,ConfiguracionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertConfiguracion(session *mgo.Session, j Configuracion) {
	c := db.Cursor(session,ConfiguracionCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllConfiguracions(session *mgo.Session) []Configuracion {
	c := db.Cursor(session,ConfiguracionCollection)
	defer session.Close()
    fmt.Println("Getting all configuracions")
	var configuracions []Configuracion
	err := c.Find(bson.M{}).All(&configuracions)
	if err != nil {
		fmt.Println(err)
	}
	return configuracions
}

func GetConfiguracionById(session *mgo.Session,id string) ([]Configuracion,error) {
	c := db.Cursor(session, ConfiguracionCollection)
	defer session.Close()
	var configuracions []Configuracion
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&configuracions)
	return configuracions,err
}

func DeleteConfiguracionById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, ConfiguracionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}