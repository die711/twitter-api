package main

import (
	"api/bd"
	"api/handlers"
	"log"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion ala BD")
		return
	}
	handlers.Manejadores()
}
