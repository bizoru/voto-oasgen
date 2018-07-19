package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const PartidoCollection = "partido"

type Partido struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
}


func UpdatePartido(session *mgo.Session, j Partido, id string) error{
	c := db.Cursor(session,PartidoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertPartido(session *mgo.Session, j Partido) {
	c := db.Cursor(session,PartidoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllPartidos(session *mgo.Session) []Partido {
	c := db.Cursor(session,PartidoCollection)
	defer session.Close()
    fmt.Println("Getting all partidos")
	var partidos []Partido
	err := c.Find(bson.M{}).All(&partidos)
	if err != nil {
		fmt.Println(err)
	}
	return partidos
}

func GetPartidoById(session *mgo.Session,id string) ([]Partido,error) {
	c := db.Cursor(session, PartidoCollection)
	defer session.Close()
	var partidos []Partido
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&partidos)
	return partidos,err
}

func DeletePartidoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, PartidoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}