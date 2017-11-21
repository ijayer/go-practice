package main

import (
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type product struct {
	PID   int     `json:"pid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) findOne(id int, session *mgo.Session) (product, error) {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(product{}).String())))

	var result product
	err := co.Find(bson.M{"pid": id}).One(&result)
	return result, err
}

func (p *product) update(condition bson.M, update bson.M, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(product{}).String())))
	return co.Update(condition, bson.M{"$set": update})
}

func (p *product) delete(condition bson.M, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(product{}).String())))
	return co.Remove(condition)
}

func (p *product) create(docs interface{}, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(product{}).String())))
	return co.Insert(docs)
}

func find(condition bson.M, session *mgo.Session, offset, limit int) ([]product, error) {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(product{}).String())))

	var v []product
	err := co.Find(condition).Skip(offset).Limit(limit).All(&v)
	return v, err
}
