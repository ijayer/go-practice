/*
 * 说明：mgo crud in go
 * 作者：zhe
 * 时间：2018-03-29 19:25
 * 更新：
 */

package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gedex/inflector"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db = "demo"

type UserStorage struct {
	User    User
	Session *mgo.Session
}

type User struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Account  string        `json:"account"`
	Password string        `json:"password"`
	Friends  []string      `json:"friends"`
	Address  Address       `json:"address"`
}

type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Remark   string `json:"remark"`
}

func (s *UserStorage) Create() error {
	session := s.Session.Copy()
	defer session.Close()

	db := session.DB(db)
	co := db.C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(s.User).Name())))
	co.EnsureIndexKey(s.User.Account)

	return co.Insert(s.User)
}

func (s *UserStorage) Upsert() error {
	session := s.Session.Copy()
	defer session.Close()

	db := session.DB(db)
	co := db.C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(s.User).Name())))

	selector := bson.M{"account": s.User.Account}
	update := &User{
		Name:     "MongoDB",
		Account:  "MongoDB",
		Password: "xxx xxx",
		Friends:  []string{"she", "he"},
		Address: Address{
			Province: "xx",
			City:     "xx",
			District: "xx",
			Remark:   "xx",
		},
	}

	_, err := co.Upsert(selector, update)
	return err
}

func (s *UserStorage) Delete() error {
	session := s.Session.Copy()
	defer session.Close()

	db := session.DB(db)
	co := db.C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(s.User).Name())))

	return co.RemoveId(s.User.Id)
}

func (s *UserStorage) Update() error {
	session := s.Session.Copy()
	defer session.Close()

	db := session.DB(db)
	co := db.C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(s.User).Name())))

	selector := bson.M{"account": s.User.Account}
	update := bson.M{"name": "golang"}

	return co.Update(selector, bson.M{"$set": update})
}

func (s *UserStorage) Find() (*User, error) {
	session := s.Session.Copy()
	defer session.Close()

	db := session.DB(db)
	co := db.C(inflector.Pluralize(strings.ToLower(reflect.TypeOf(s.User).Name())))

	user := new(User)
	err := co.Find(bson.M{"account": s.User.Account}).One(user)
	return user, err
}

func main() {
	session := InitDB()

	user := User{
		Id:       bson.NewObjectId(),
		Name:     "Gopher",
		Account:  "Gopher",
		Password: "******",
		Friends:  []string{"you", "me"},
		Address: Address{
			Province: "zj",
			City:     "hz",
			District: "gs",
			Remark:   "soho",
		},
	}

	us := UserStorage{User: user, Session: session}

	//err := us.Create()
	//if err != nil {
	//	fmt.Println("Error: create failed:", err.Error())
	//}
	//
	//u, err := us.Find()
	//if err != nil {
	//	fmt.Println("Error: find failed:", err.Error())
	//}
	//fmt.Printf("find result: %+v\n", *u)
	//
	//err = us.Update()
	//if err != nil {
	//	fmt.Println("Error: update failed:", err.Error())
	//}

	err := us.Upsert()
	if err != nil {
		fmt.Println("Error: upsert failed:", err.Error())
	}
}

func InitDB() *mgo.Session {
	var err error
	var session *mgo.Session

	info := &mgo.DialInfo{}
	info.Timeout = 60 * time.Second
	info.Database = db
	info.Addrs = []string{"127.0.0.1"}

	session, err = mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic(单调的) behavior(行为).
	session.SetMode(mgo.Monotonic, true)

	return session
}
