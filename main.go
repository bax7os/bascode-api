package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fs := http.FileServer(http.Dir("./uploads/fotos-perfil"))
	http.Handle("/fotos/", http.StripPrefix("/fotos/", fs))

	// Configuração para servir outros arquivos estáticos (CSS, JS, imagens padrão)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
