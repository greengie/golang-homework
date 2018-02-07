package main

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

type Product struct {
  ID bson.ObjectId `bson:"_id" json:"id"`
  ProductID string  `bson:"ProductID" json:"ProductID"`
  ProductName string `bson:"ProductName" json:"ProductName"`
  SupplierID string `bson:"SupplierID" json:"SupplierID"`
  CategoryID string `bson:"CategoryID" json:"CategoryID"`
  QuantityPerUnit string `bson:"QuantityPerUnit" json:"QuantityPerUnit"`
  UnitPrice float64 `bson:"UnitPrice" json:"UnitPrice"`
  UnitsInStock int64 `bson:"UnitsInStock" json:"UnitsInStock"`
  UnitsOnOrder int64 `bson:"UnitsOnOrder" json:"UnitsOnOrder"`
  ReorderLevel int64 `bson:"ReorderLevel" json:"ReorderLevel"`
  Discontinued int64 `bson:"Discontinued" json:"Discontinued"`
  Time time.Time `bson:"Timestamp" json:"Timestamp"`
}

type Order struct {
  ID bson.ObjectId `bson:"_id" json:"id"`
  OrderID string `bson:"OrderID" json:"OrderID"`
}
