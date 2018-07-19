package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const SufraganteCollection = "sufragante"

type Sufragante struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Sufragante []Votante `json:"sufragantes"`
  Evento []Evento `json:"eventos"`
}


func UpdateSufragante(session *mgo.Session, j Sufragante, id string) error{
	c := db.Cursor(session,SufraganteCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertSufragante(session *mgo.Session, j Sufragante) {
	c := db.Cursor(session,SufraganteCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllSufragantes(session *mgo.Session) []Sufragante {
	c := db.Cursor(session,SufraganteCollection)
	defer session.Close()
    fmt.Println("Getting all sufragantes")
	var sufragantes []Sufragante
	err := c.Find(bson.M{}).All(&sufragantes)
	if err != nil {
		fmt.Println(err)
	}
	return sufragantes
}

func GetSufraganteById(session *mgo.Session,id string) ([]Sufragante,error) {
	c := db.Cursor(session, SufraganteCollection)
	defer session.Close()
	var sufragantes []Sufragante
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&sufragantes)
	return sufragantes,err
}

func DeleteSufraganteById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, SufraganteCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}