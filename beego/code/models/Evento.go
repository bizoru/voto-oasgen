package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
  "time"
)

const EventoCollection = "evento"

type Evento struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Tipo string `json:"tipo"`
  Modulo string `json:"modulo"`
  Usuario string `json:"usuario"`
  Fecha time.Time `json:"fecha"`
  Ip string `json:"ip"`
}


func UpdateEvento(session *mgo.Session, j Evento, id string) error{
	c := db.Cursor(session,EventoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertEvento(session *mgo.Session, j Evento) {
	c := db.Cursor(session,EventoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllEventos(session *mgo.Session) []Evento {
	c := db.Cursor(session,EventoCollection)
	defer session.Close()
    fmt.Println("Getting all eventos")
	var eventos []Evento
	err := c.Find(bson.M{}).All(&eventos)
	if err != nil {
		fmt.Println(err)
	}
	return eventos
}

func GetEventoById(session *mgo.Session,id string) ([]Evento,error) {
	c := db.Cursor(session, EventoCollection)
	defer session.Close()
	var eventos []Evento
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&eventos)
	return eventos,err
}

func DeleteEventoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, EventoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}