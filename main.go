package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"mongo-golang/controllers"
)

func main(){

	r := httprouter.New()
	uc := controllers.NewCategoryController(getSession())
	r.GET("/category/:id", uc.GetCategory)
	r.POST("/category", uc.CreateCategory)
	r.DELETE("/category/:id", uc.DeleteCategory)
	http.ListenAndServe("localhost:9000", r)
}


func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}