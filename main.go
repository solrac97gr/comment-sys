package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/solrac97gr/comments/migration"
	"github.com/solrac97gr/comments/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzó la migración")
		migration.Migrate()
		log.Println("Finalizó la migración")
	}
	// Init the routes
	router := routes.InitRoutes()
	// Init the middlewares
	n := negroni.Classic()
	n.UseHandler(router)
	// Init the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecución del programa")
}
