package main

import(
    "net/http"
    "encoding/json"

    // log "github.com/Sirupsen/logrus"
    "gopkg.in/mgo.v2/bson"
    "time"
)

func CreateProductEndPoint(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
  var product Product
  if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		JSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = bson.NewObjectId()
  product.Time = time.Now()
	dBcontroller := RP.Request(r)
	if err := dBcontroller.InsertProduct(product); err != nil {
		JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	JSONResponse(w, http.StatusCreated, product)

}

func UpdateOrderEndPoint(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		JSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	dBcontroller := RP.Request(r)
	if err := dBcontroller.UpdateOrder(order); err != nil {
		JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	JSONResponse(w, http.StatusOK, map[string]string{"result": "success"})
}
