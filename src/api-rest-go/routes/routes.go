package routes

import (
	"log"
	"net/http"

	"github.com/RafaelSanzio0/GO/go-rest-api/controllers"
	"github.com/RafaelSanzio0/GO/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home) // quando cair no endereço barra vazio joga pra função home atender
	r.HandleFunc("/api/personalidades", controllers.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadesById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
