package main

import (
	"log"
	"net/http"
	"problem3_irdeto/internal/api"
	"problem3_irdeto/internal/server"
)

func main(){
	r := server.SetServer()
	c := api.NewAPIClient()
	h := api.NewAPIHandler(c)
	h.RegisterRoute(r)
	// set TCP network address
	log.Fatal(http.ListenAndServe(":8000", r))
}
