package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const TarjetonCollection = "tarjeton"

type Tarjeton struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Candidato []Candidato `json:"candidatos"`
  Eleccion []Eleccion `json:"eleccions"`
  Configuracion []Configuracion `json:"configuracions"`
}


func UpdateTarjeton(session *mgo.Session, j Tarjeton, id string) error{
	c := db.Cursor(session,TarjetonCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertTarjeton(session *mgo.Session, j Tarjeton) {
	c := db.Cursor(session,TarjetonCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllTarjetons(session *mgo.Session) []Tarjeton {
	c := db.Cursor(session,TarjetonCollection)
	defer session.Close()
    fmt.Println("Getting all tarjetons")
	var tarjetons []Tarjeton
	err := c.Find(bson.M{}).All(&tarjetons)
	if err != nil {
		fmt.Println(err)
	}
	return tarjetons
}

func GetTarjetonById(session *mgo.Session,id string) ([]Tarjeton,error) {
	c := db.Cursor(session, TarjetonCollection)
	defer session.Close()
	var tarjetons []Tarjeton
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&tarjetons)
	return tarjetons,err
}

func DeleteTarjetonById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, TarjetonCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}