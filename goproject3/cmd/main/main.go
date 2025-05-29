package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//_ "github.com/jinzhu/gormdialects/mysql"
	"github.com/AnantAgarwal1008/Golang-Minor-Projects/pkg/routes"
)
//bookstore
func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}