package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const EleccionCollection = "eleccion"

type Eleccion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
  Votante []Votante `json:"votantes"`
  Ponderacion []Ponderacion `json:"ponderacions"`
  Responsable []Usuario `json:"responsables"`
  Habilitado bool `json:"habilitado"`
}


func UpdateEleccion(session *mgo.Session, j Eleccion, id string) error{
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertEleccion(session *mgo.Session, j Eleccion) {
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllEleccions(session *mgo.Session) []Eleccion {
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
    fmt.Println("Getting all eleccions")
	var eleccions []Eleccion
	err := c.Find(bson.M{}).All(&eleccions)
	if err != nil {
		fmt.Println(err)
	}
	return eleccions
}

func GetEleccionById(session *mgo.Session,id string) ([]Eleccion,error) {
	c := db.Cursor(session, EleccionCollection)
	defer session.Close()
	var eleccions []Eleccion
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&eleccions)
	return eleccions,err
}

func DeleteEleccionById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, EleccionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}