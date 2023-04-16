package main

import (
	_ "bytes"
	"fmt"
	"log"
	"net/http"

	_ "github.com/godror/godror"
	"github.com/rs/cors"
)

func main() {
	iniciarServidor("127.0.0.0", "3000")
}

func iniciarServidor(ip string, puerto string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})
	router := NewRouter()

	fmt.Println("El servidor esta corriendo en http://" + ip + ":" + puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, c.Handler(router)))
}
