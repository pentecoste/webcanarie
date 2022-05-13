/*
 * main.go
 *
 * Codice principale del programma.
 *
 * Copyright (c) 2022 Davide Vendramin <davidevendramin5@gmail.com>
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pentecoste/webcanarie/internal/auth"
	"github.com/pentecoste/webcanarie/internal/config"
	"github.com/pentecoste/webcanarie/internal/handlers"
	"github.com/pentecoste/webcanarie/internal/db"
)

var Version string

func main() {
	if err := config.LoadConfig("config/config.toml"); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.InizializzaDB(); err != nil {
		fmt.Println(err)
		return
	}

	defer db.ChiudiDB()

	auth.InitializeSigning()

	fmt.Println("WebCanarie versione: " + Version)

	// Imposta la versione nei package che lo richiedono
	handlers.Version = Version

	mux := http.NewServeMux()

	// I pattern che finiscono per '/' comprendono anche i sottopercorsi.
	// Sono valutati 'a partire dal più specifico', quindi '/' sarà
	// sempre l'ultimo.
	mux.HandleFunc("/", handlers.HandleRootOr404)
	mux.HandleFunc("/apartment/", handlers.HandleApartment)
	mux.HandleFunc("/contacts", handlers.HandleContacts)
	mux.HandleFunc("/availability", handlers.HandleAvailability)
	mux.HandleFunc("/whereis", handlers.HandleWhereIs)

	// File server per servire direttamente i contenuti statici.
	fileserver := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	srvAddress := config.Config.Generale.Porta
	srv := &http.Server{
		Addr:    srvAddress,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
