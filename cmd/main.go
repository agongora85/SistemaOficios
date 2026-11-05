package main

import (
	"fmt"
	"sistema-oficios/internal/database"
	"sistema-oficios/internal/handlers"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.GET(
		"/oficios",
		handlers.ObtenerOficios,
	)

	r.Run(":8080")

	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	// Declaración de variables
	// var firstName, lastName string
	// var age int
	var (
		firstName, lastName string
		age                 int
	)

	var calle, cp = "78", "97320"
	tel, ext := "985555", "125"
	firstName = "Alex"
	lastName = "Alex"
	age = 27

	fmt.Println(firstName, lastName, age)
	fmt.Println(calle, cp)
	fmt.Println(tel, ext)
}
