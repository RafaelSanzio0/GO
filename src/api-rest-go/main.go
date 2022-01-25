package main

import (
	"fmt"

	"github.com/RafaelSanzio0/GO/go-rest-api/database"
	"github.com/RafaelSanzio0/GO/go-rest-api/routes"
)

func main() {
	database.ConnectDataBase()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest() // fica ouvindo a handlefunc
}
