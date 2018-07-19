package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const HistoriaCollection = "historia"

type Historia struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Jornada string `json:"jornada"`
  Resultado []Eleccion `json:"resultados"`
}


func UpdateHistoria(session *mgo.Session, j Historia, id string) error{
	c := db.Cursor(session,HistoriaCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertHistoria(session *mgo.Session, j Historia) {
	c := db.Cursor(session,HistoriaCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllHistorias(session *mgo.Session) []Historia {
	c := db.Cursor(session,HistoriaCollection)
	defer session.Close()
    fmt.Println("Getting all historias")
	var historias []Historia
	err := c.Find(bson.M{}).All(&historias)
	if err != nil {
		fmt.Println(err)
	}
	return historias
}

func GetHistoriaById(session *mgo.Session,id string) ([]Historia,error) {
	c := db.Cursor(session, HistoriaCollection)
	defer session.Close()
	var historias []Historia
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&historias)
	return historias,err
}

func DeleteHistoriaById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, HistoriaCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}