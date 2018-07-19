package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)

const CertificadoCollection = "certificado"

type Certificado struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Proceso []Proceso `json:"procesos"`
  Eleccion []Eleccion `json:"eleccions"`
  Sufragante []Sufragante `json:"sufragantes"`
}


func UpdateCertificado(session *mgo.Session, j Certificado, id string) error{
	c := db.Cursor(session,CertificadoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertCertificado(session *mgo.Session, j Certificado) {
	c := db.Cursor(session,CertificadoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllCertificados(session *mgo.Session) []Certificado {
	c := db.Cursor(session,CertificadoCollection)
	defer session.Close()
    fmt.Println("Getting all certificados")
	var certificados []Certificado
	err := c.Find(bson.M{}).All(&certificados)
	if err != nil {
		fmt.Println(err)
	}
	return certificados
}

func GetCertificadoById(session *mgo.Session,id string) ([]Certificado,error) {
	c := db.Cursor(session, CertificadoCollection)
	defer session.Close()
	var certificados []Certificado
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&certificados)
	return certificados,err
}

func DeleteCertificadoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, CertificadoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}