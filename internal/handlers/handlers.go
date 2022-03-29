/*
 * handlers.go
 *
 * Package per gestire le diverse pagine ed i relativi template.
 *
 * Copyright (c) 2020 Davide Vendramin <davidevendramin5@gmail.com>
 */

// Package per gestire le diverse pagine ed i relativi template.
package handlers

import (
	_ "encoding/json"
	_ "fmt"
	_ "github.com/pentecoste/webcanarie/internal/auth"
	_ "github.com/pentecoste/webcanarie/internal/config"
	_ "github.com/pentecoste/webcanarie/internal/hash"
	"net/http"
	_ "net/url"
	_ "strconv"
	_ "strings"
	"text/template"
)

const templatesDir = "web/template"

var Version string

type CommonValues struct {
	Version string
}

// viene inizializzato nel momento in cui viene importato il package
var templates = template.Must(template.ParseFiles(
	templatesDir+"/index.html",
	templatesDir+"/appartamento.html",
))

// Handler per qualunque percorso diverso da tutti gli altri percorsi riconosciuti.
// Caso particolare è la homepage (/); per ogni altro restituisce 404.
func HandleRootOr404(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	HandleHome(w, r)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func handleApartment(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "appartamento.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}
