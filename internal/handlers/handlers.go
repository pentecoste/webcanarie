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
	_ "github.com/pentecoste/webcanarie/internal/auth"
	_ "github.com/pentecoste/webcanarie/internal/config"
	"github.com/pentecoste/webcanarie/internal/db"
	_ "github.com/pentecoste/webcanarie/internal/hash"
	_ "io/ioutil"
	"net/http"
	"time"
	_ "net/url"
	"strconv"
	"strings"
	"text/template"
)

const templatesDir = "web/template"

var Version string

type CommonValues struct {
	Version string
}

type Giorno struct {
	Giorno		byte
	Prenotato	bool
}

type Calendario struct {
	Mese	string
	Anno	uint32
	Giorni	[]Giorno
	Delay	[]byte
	Buffer  []byte
}

type RecensioneElaborata struct {
	Descrizione	string
	Data		string
	Persone		string
	Icone		[]string
}

// viene inizializzato nel momento in cui viene importato il package
var templates = template.Must(template.ParseFiles(
	templatesDir+"/index.html",
	templatesDir+"/appartamento.html",
	templatesDir+"/contacts.html",
	templatesDir+"/availability.html",
	templatesDir+"/whereis.html",
	templatesDir+"/recensioni.html",
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
	img_app, err := db.GetImmaginiAppartamento()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img_isl, err := db.GetImmaginiIsola()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "index.html", struct {
		ImgApp	    []db.Immagine
		ImgIsl	    []db.Immagine
		Values      CommonValues
	}{img_app, img_isl, CommonValues{Version}})
}

func HandleApartment(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/apartment/")
	idParsed, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Redirect(w, r, "/apartment/1", http.StatusSeeOther)
		return
	}

	idStanza := uint32(idParsed)
	imgs, err := db.GetImmaginiByStanza(idStanza)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if imgs == nil {
		http.Redirect(w, r, "/apartment/1", http.StatusSeeOther)
		return
	}

	stanze, err := db.GetStanze()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stanza, err := db.GetStanza(idStanza)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*isRoom := make([]byte, len(stanze))
	isRoom[idStanza] = 1*/

	templates.ExecuteTemplate(w, "appartamento.html", struct {
		Immagini    []db.Immagine
		Stanze	    []db.Stanza
		Stanza	    db.Stanza
		IdStanza    uint32
		Values      CommonValues
	}{imgs, stanze, stanza, idStanza-1, CommonValues{Version}})
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contacts.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func HandleAvailability(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	mesi := map[string]string{
		"January": "Gennaio",
		"February": "Febbraio",
		"March": "Marzo",
		"April": "Aprile",
		"May": "Maggio",
		"June": "Giugno",
		"July": "Luglio",
		"August": "Agosto",
		"September": "Settembre",
		"October": "Ottobre",
		"November": "Novembre",
		"December": "Dicembre",
	}

	var months[12] Calendario;
	newTime := now
	for i:=0; i<12; i++{
		months[i].Mese = mesi[newTime.Month().String()]
		months[i].Anno = uint32(newTime.Year())
		weekDay := byte(int(time.Date(newTime.Year(), newTime.Month(), 1, 0, 0, 0, 0, time.UTC).Weekday()))
		if weekDay != 0 {
			weekDay -= 1
		} else {
			weekDay = 6
		}
		months[i].Delay = make([]byte, weekDay)
		newTime = newTime.AddDate(0,1,0)
		daysLen := time.Date(newTime.Year(), newTime.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
		months[i].Buffer = make([]byte, 7-(weekDay+byte(daysLen))%7)
	}

	prenotazioni, err := db.GetLastPrenotazioni()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTime = now
	for i:=0; i<12; i++{
		daysLen := time.Date(newTime.Year(), newTime.AddDate(0,1,0).Month(), 0, 0, 0, 0, 0, time.UTC).Day()
		months[i].Giorni = make([]Giorno, daysLen)
		for j:=0; j<daysLen; j++{
			months[i].Giorni[j].Giorno = byte(j+1)
			booked := false
			unix_ts := time.Date(newTime.Year(), newTime.Month(), 1, 0, 0, 0, 0, time.UTC).Unix() + int64(j*86400)
			for _, prenotazione := range prenotazioni{
				if (unix_ts >= prenotazione.Inizio.Unix() && unix_ts < prenotazione.Fine.Unix()){
					booked = true
					break
				}
			}
			months[i].Giorni[j].Prenotato = booked
		}
		newTime = newTime.AddDate(0,1,0)
	}

	templates.ExecuteTemplate(w, "availability.html", struct {
		Months	    [12]Calendario
		Values      CommonValues
	}{months, CommonValues{Version}})
}

func HandleWhereIs(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "whereis.html", struct {
		Values      CommonValues
	}{CommonValues{Version}})
}

func HandleFeedbacks(w http.ResponseWriter, r *http.Request) {
	recensioni, err := db.GetRecensioni()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mesi := map[string]string{
		"January": "gennaio",
		"February": "febbraio",
		"March": "marzo",
		"April": "aprile",
		"May": "maggio",
		"June": "giugno",
		"July": "luglio",
		"August": "agosto",
		"September": "settembre",
		"October": "ottobre",
		"November": "novembre",
		"December": "dicembre",
	}

	reviews := make([]RecensioneElaborata, len(recensioni))
	for i, recensione := range recensioni {
		reviews[i].Descrizione = recensione.Descrizione
		reviews[i].Data = mesi[recensione.Data.Month().String()]
		reviews[i].Data += " " + strconv.Itoa(recensione.Data.Year())
		reviews[i].Persone = recensione.Persone
		for _, c := range recensione.Icone {
			switch c {
				case 'M':
					reviews[i].Icone = append(reviews[i].Icone, "adult_male")
				case 'F':
					reviews[i].Icone = append(reviews[i].Icone, "adult_female")
				case 'm':
					reviews[i].Icone = append(reviews[i].Icone, "child_male")
				default:
					reviews[i].Icone = append(reviews[i].Icone, "child_female")
			}
		}
	}

	templates.ExecuteTemplate(w, "recensioni.html", struct {
		Recensioni  []RecensioneElaborata
		Values      CommonValues
	}{reviews, CommonValues{Version}})
}
