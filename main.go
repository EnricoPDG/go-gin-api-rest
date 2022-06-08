package main

import (
	"github.com/EnricoPDG/go-gin-api-rest/models"
	"github.com/EnricoPDG/go-gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Enrico", CPF: "123.456.789-10", RG: "11111111"},
		{Nome: "Nath", CPF: "123.456.333-30", RG: "22222222"},
	}
	routes.HandleRequests()
}
