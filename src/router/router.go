package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar returns a router with configured routes
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
