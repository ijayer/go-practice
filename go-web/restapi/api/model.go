package api

import (
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Product model
type Product struct {
	PID   int     `json:"pid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// findOne
func (p *Product) findOne(id int, session *mgo.Session) (Product, error) {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(Product{}).String())))

	var result Product
	err := co.Find(bson.M{"pid": id}).One(&result)
	return result, err
}

// update
func (p *Product) update(condition bson.M, update bson.M, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(Product{}).String())))
	return co.Update(condition, bson.M{"$set": update})
}

// delete
func (p *Product) delete(condition bson.M, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(Product{}).String())))
	return co.Remove(condition)
}

// create
func (p *Product) create(docs interface{}, session *mgo.Session) error {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(Product{}).String())))
	return co.Insert(docs)
}

// find
func find(condition bson.M, session *mgo.Session, offset, limit int) ([]Product, error) {
	session = session.Copy()
	defer session.Close()
	co := session.DB(DBName).C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(Product{}).String())))

	var v []Product
	err := co.Find(condition).Skip(offset).Limit(limit).All(&v)
	return v, err
}
