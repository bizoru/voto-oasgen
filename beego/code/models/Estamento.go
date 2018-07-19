package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const EstamentoCollection = "estamento"

type Estamento struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
}


func UpdateEstamento(session *mgo.Session, j Estamento, id string) error{
	c := db.Cursor(session,EstamentoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertEstamento(session *mgo.Session, j Estamento) {
	c := db.Cursor(session,EstamentoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllEstamentos(session *mgo.Session) []Estamento {
	c := db.Cursor(session,EstamentoCollection)
	defer session.Close()
    fmt.Println("Getting all estamentos")
	var estamentos []Estamento
	err := c.Find(bson.M{}).All(&estamentos)
	if err != nil {
		fmt.Println(err)
	}
	return estamentos
}

func GetEstamentoById(session *mgo.Session,id string) ([]Estamento,error) {
	c := db.Cursor(session, EstamentoCollection)
	defer session.Close()
	var estamentos []Estamento
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&estamentos)
	return estamentos,err
}

func DeleteEstamentoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, EstamentoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}