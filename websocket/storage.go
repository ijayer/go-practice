package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

type UserStorage struct {
	CollectionName  string
	Session         *mgo.Session
}

func (s *UserResource) Insert() {
	fmt.Println("##_______________Into Insert")
}

func (s *UserResource) Login() {
	fmt.Println("##_______________Into Login")
}

func (s *UserResource) Logout() {
	fmt.Println("##_______________Into Logout")
}