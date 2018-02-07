package main

import (
  "net/http"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  // "fmt"
)

type Repo struct {
  mongoSession *mgo.Session
}

var RP Repo

func (r Repo) Request(req *http.Request) *Repo {
  return &Repo{mongoSession: getMongoSession()}
}

func (r Repo) InsertProduct(product Product) error {
  session := r.mongoSession.Clone()
  defer session.Close()

  err := session.DB("northwind").C("products").Insert(&product)
  return err
}

func (r Repo) UpdateOrder(order Order) error {
  session := r.mongoSession.Clone()
  defer session.Close()

  selector := bson.M{"OrderID": &order.OrderID}
  update := bson.M{"$set": bson.M{ "Discount": 0.3}}

  _, err := session.DB("northwind").C("order-details").UpdateAll(selector, update)
  return err
}
