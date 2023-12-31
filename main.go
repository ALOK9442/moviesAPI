package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ALOK9442/mongoAPI/router"
)

func main() {
	fmt.Println("welcome to movies api")
	fmt.Println("server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000")
}
