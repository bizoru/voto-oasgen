package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const CandidatoCollection = "candidato"

type Candidato struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Nombre string `json:"nombre"`
  Apellidos string `json:"apellidos"`
  Correo string `json:"correo"`
  Dependencia string `json:"dependencia"`
  Partido string `json:"partido"`
  Foto string `json:"foto"`
}


func UpdateCandidato(session *mgo.Session, j Candidato, id string) error{
	c := db.Cursor(session,CandidatoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertCandidato(session *mgo.Session, j Candidato) {
	c := db.Cursor(session,CandidatoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllCandidatos(session *mgo.Session) []Candidato {
	c := db.Cursor(session,CandidatoCollection)
	defer session.Close()
    fmt.Println("Getting all candidatos")
	var candidatos []Candidato
	err := c.Find(bson.M{}).All(&candidatos)
	if err != nil {
		fmt.Println(err)
	}
	return candidatos
}

func GetCandidatoById(session *mgo.Session,id string) ([]Candidato,error) {
	c := db.Cursor(session, CandidatoCollection)
	defer session.Close()
	var candidatos []Candidato
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&candidatos)
	return candidatos,err
}

func DeleteCandidatoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, CandidatoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}