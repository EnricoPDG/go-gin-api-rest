package main

import (
	"github.com/EnricoPDG/go-gin-api-rest/database"
	"github.com/EnricoPDG/go-gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
