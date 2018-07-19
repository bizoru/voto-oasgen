package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
  "time"
)

const ProcesoCollection = "proceso"

type Proceso struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Eleccion []Eleccion `json:"eleccions"`
  Nombre string `json:"nombre"`
  Fecha_inicio time.Time `json:"fecha_inicio"`
  Fecha_final time.Time `json:"fecha_final"`
}


func UpdateProceso(session *mgo.Session, j Proceso, id string) error{
	c := db.Cursor(session,ProcesoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertProceso(session *mgo.Session, j Proceso) {
	c := db.Cursor(session,ProcesoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllProcesos(session *mgo.Session) []Proceso {
	c := db.Cursor(session,ProcesoCollection)
	defer session.Close()
    fmt.Println("Getting all procesos")
	var procesos []Proceso
	err := c.Find(bson.M{}).All(&procesos)
	if err != nil {
		fmt.Println(err)
	}
	return procesos
}

func GetProcesoById(session *mgo.Session,id string) ([]Proceso,error) {
	c := db.Cursor(session, ProcesoCollection)
	defer session.Close()
	var procesos []Proceso
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&procesos)
	return procesos,err
}

func DeleteProcesoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, ProcesoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}