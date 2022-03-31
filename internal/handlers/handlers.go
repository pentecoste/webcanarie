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
	_ "io/ioutil"
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
	templatesDir+"/contacts.html",
	templatesDir+"/whereis.html",
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
	/*apartment_imgs, err := ioutil.ReadDir("web/static/img/apartment/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isle_imgs, err := ioutil.ReadDir("web/static/img/isle/")
        if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	count := 0
	for _, img := range apartment_imgs {
		if !(img.IsDir()) {
			count++
		}
	}

	apartment_names := make([]string, count)
	for _, img range apartment_imgs {
		if !(img.IsDir()) {
			apartment_names = append(apartment_names, img.Name())
		}
	}

	count = 0
	for _, img := range isle_imgs {
		if !(img.IsDir()) {
			count++
		}
	}

	isle_names := make([]string, count)
	for _, img range isle_imgs {
		if !(img.IsDir()) {
			isle_names = append(isle_names, img.Name())
		}
	}

	templates.ExecuteTemplate(w, "index.html", struct {
		AppImgs	    []string
		IsleImgs    []string
		Values      CommonValues
	}{apartment_names, isle_names, CommonValues{Version}})*/

	templates.ExecuteTemplate(w, "index.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func HandleApartment(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "appartamento.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contacts.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func HandleWhereIs(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "whereis.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}
