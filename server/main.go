package main

import (
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }
	//init db
	mysql.DatabaseInit()

	//Run Migration
	database.RunMigration()

	r := mux.NewRouter()

	// routes.Route(router.PathPrefix("api/v1").Subrouter())
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())
	// Setup allowed Header, Method, and Origin for CORS on this below code ...
var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})
r.PathPrefix("/upload").Handler(http.StripPrefix("/upload/",http.FileServer(http.Dir("./upload"))))

var port = "5000"
fmt.Println("server running localhost:"+port)

http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
	
}