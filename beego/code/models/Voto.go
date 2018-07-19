package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const VotoCollection = "voto"

type Voto struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Candidato []Candidato `json:"candidatos"`
  Evento []Evento `json:"eventos"`
}


func UpdateVoto(session *mgo.Session, j Voto, id string) error{
	c := db.Cursor(session,VotoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertVoto(session *mgo.Session, j Voto) {
	c := db.Cursor(session,VotoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllVotos(session *mgo.Session) []Voto {
	c := db.Cursor(session,VotoCollection)
	defer session.Close()
    fmt.Println("Getting all votos")
	var votos []Voto
	err := c.Find(bson.M{}).All(&votos)
	if err != nil {
		fmt.Println(err)
	}
	return votos
}

func GetVotoById(session *mgo.Session,id string) ([]Voto,error) {
	c := db.Cursor(session, VotoCollection)
	defer session.Close()
	var votos []Voto
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&votos)
	return votos,err
}

func DeleteVotoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, VotoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}