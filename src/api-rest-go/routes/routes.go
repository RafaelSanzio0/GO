package routes

import (
	"log"
	"net/http"

	"github.com/RafaelSanzio0/GO/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home) // quando cair no endereço barra vazio joga pra função home atender
	r.HandleFunc("/api/personalidades", controllers.AllPersonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadesById)

	log.Fatal(http.ListenAndServe(":8000", r))
}
