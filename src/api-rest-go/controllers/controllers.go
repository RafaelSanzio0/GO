package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RafaelSanzio0/GO/go-rest-api/database"
	"github.com/RafaelSanzio0/GO/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade

	database.DB.Find(&p) // encontra no banco a personalidade

	json.NewEncoder(w).Encode(p) // quero ter um novo encoder do tipo JSON e encondar todas as minhas personalidades
}

func PersonalidadesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}
