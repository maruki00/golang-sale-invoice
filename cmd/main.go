package main

import (
	routes "delivery/golang_salesInvoice/UserInterface/Routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterCategoryRouter(router)

	http.Handle("/", router)
	fmt.Println("Server Started On http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
