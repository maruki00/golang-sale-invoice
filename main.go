package main

import (
	routes "delivery/golang_salesInvoice/UserInterface/Routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routes.RegisterCategoryRouter(route)

	http.Handle("/", route)
	fmt.Println("Server Started On http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
