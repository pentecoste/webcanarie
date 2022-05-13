package db

import (
	"database/sql"
	"fmt"
	"github.com/pentecoste/webcanarie/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//Tabelle del database
type Immagine struct {
	Codice		uint32
	Percorso        string
	Descrizione     string
	Stanza		string
}

type Stanza struct {
	Codice		uint32
	Nome		string
	Dotazioni	string
}

type Prenotazione struct {
	Codice		uint32
	Inizio		time.Time
	Fine		time.Time
}

var db_Connection *sql.DB

//Funzione per inizializzare il database
func InizializzaDB() (err error) {
	db_Connection, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", config.Config.SQL.Username, config.Config.SQL.Password, config.Config.SQL.Indirizzo, config.Config.SQL.Database))
	return
}

//Funzione per chiudere il database
func ChiudiDB() {
	db_Connection.Close()
}

func GetImmagine(cod uint32) (Immagine, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err := db_Connection.Ping(); err != nil {
		return Immagine{}, err
	}

	//Salvo la query che eseguirà l'sql in una variabile stringa
	q := `SELECT Immagine.Codice,Immagine.Percorso,Immagine.Descrizione,Stanza.Nome FROM Immagine,Stanza WHERE Immagine.Stanza = Stanza.Codice AND Immagine.Codice = ?`
	//Applico la query al database. Salvo i risultati in rows
	rows, err := db_Connection.Query(q, cod)
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err != nil {
		return Immagine{}, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	//Creo un libro in cui salvare il risultato della query
	var img Immagine
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		//Tramite rows.Scan() salvo i vari risultati nel libro creato in precedenza. In caso di errore ritorno un libro vuoto e l'errore
		if err := rows.Scan(&img.Codice, &img.Percorso, &img.Descrizione, &img.Stanza); err != nil {
			return Immagine{}, err
		}
	}
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err := rows.Err(); err != nil {
		return Immagine{}, err
	}

	//Returno il libro trovato e null (null sarebbe l'errore che non è avvenuto)
	return img, nil
}

func GetStanza(cod uint32) (Stanza, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err := db_Connection.Ping(); err != nil {
		return Stanza{}, err
	}

	//Salvo la query che eseguirà l'sql in una variabile stringa
	q := `SELECT Stanza.Codice,Stanza.Nome,Stanza.Dotazioni FROM Stanza WHERE Stanza.Codice = ?`
	//Applico la query al database. Salvo i risultati in rows
	rows, err := db_Connection.Query(q, cod)
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err != nil {
		return Stanza{}, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	//Creo un libro in cui salvare il risultato della query
	var stn Stanza
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		//Tramite rows.Scan() salvo i vari risultati nel libro creato in precedenza. In caso di errore ritorno un libro vuoto e l'errore
		if err := rows.Scan(&stn.Codice, &stn.Nome, &stn.Dotazioni); err != nil {
			return Stanza{}, err
		}
	}
	//Se c'è un errore, ritorna un libro vuoto e l'errore
	if err := rows.Err(); err != nil {
		return Stanza{}, err
	}

	//Returno il libro trovato e null (null sarebbe l'errore che non è avvenuto)
	return stn, nil
}

func GetImmaginiByStanza(stanza uint32) ([]Immagine, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Immagine.Codice,Immagine.Percorso,Immagine.Descrizione,Stanza.Nome
	      FROM Immagine,Stanza
		  WHERE Immagine.Stanza = Stanza.Codice AND Stanza.Codice = ?`

	rows, err := db_Connection.Query(q, stanza)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var imgs []Immagine
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Immagine
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &fabrizio.Percorso, &fabrizio.Descrizione, &fabrizio.Stanza); err != nil {
			return nil, err
		}
		//Copio la variabile temporanea nell'ultima posizione dell'array
		imgs = append(imgs, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return imgs, nil
}

func GetImmagini() ([]Immagine, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Immagine.Codice,Immagine.Percorso,Immagine.Descrizione,Stanza.Nome
	      FROM Immagine,Stanza
		  WHERE Immagine.Stanza = Stanza.Codice`

	rows, err := db_Connection.Query(q)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var imgs []Immagine
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Immagine
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &fabrizio.Percorso, &fabrizio.Descrizione, &fabrizio.Stanza); err != nil {
			return nil, err
		}
		//Copio la variabile temporanea nell'ultima posizione dell'array
		imgs = append(imgs, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return imgs, nil
}

func GetImmaginiAppartamento() ([]Immagine, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Immagine.Codice,Immagine.Percorso,Immagine.Descrizione,Stanza.Nome
	      FROM Immagine,Stanza
		  WHERE Immagine.Stanza = Stanza.Codice AND Immagine.Stanza IS NOT NULL`

	rows, err := db_Connection.Query(q)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var imgs []Immagine
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Immagine
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &fabrizio.Percorso, &fabrizio.Descrizione, &fabrizio.Stanza); err != nil {
			return nil, err
		}
		//Copio la variabile temporanea nell'ultima posizione dell'array
		imgs = append(imgs, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return imgs, nil
}

func GetImmaginiIsola() ([]Immagine, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Immagine.Codice,Immagine.Percorso,Immagine.Descrizione
	      FROM Immagine
		  WHERE Immagine.Stanza IS NULL`

	rows, err := db_Connection.Query(q)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var imgs []Immagine
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Immagine
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &fabrizio.Percorso, &fabrizio.Descrizione); err != nil {
			return nil, err
		}
		fabrizio.Stanza = ""
		//Copio la variabile temporanea nell'ultima posizione dell'array
		imgs = append(imgs, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return imgs, nil
}

func GetStanze() ([]Stanza, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Stanza.Codice,Stanza.Nome,Stanza.Dotazioni
	      FROM Stanza`

	rows, err := db_Connection.Query(q)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var stns []Stanza
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Stanza
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &fabrizio.Nome, &fabrizio.Dotazioni); err != nil {
			return nil, err
		}
		//Copio la variabile temporanea nell'ultima posizione dell'array
		stns = append(stns, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stns, nil
}

func GetLastPrenotazioni() ([]Prenotazione, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return nil, err
	}

	now := time.Now()

	beginning := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	ending := beginning.AddDate(1, 0, 0)

	//Esamino tutti i casi possibili di richiesta, scegliendo la query giusta per ogni situazione possibile
	q := `SELECT Prenotazione.Codice,Prenotazione.Inizio,Prenotazione.Fine
	      FROM Prenotazione
	      WHERE Prenotazione.Inizio > ? OR Prenotazione.Fine < ?`

	rows, err := db_Connection.Query(q, beginning.Unix(), ending.Unix())
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return nil, err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	var prens []Prenotazione
	//Rows.Next() scorre tutte le righe trovate dalla query returnando true. Quando le finisce returna false
	for rows.Next() {
		var fabrizio Prenotazione
		var inizio int64
		var fine int64
		//Tramite rows.Scan() salvo i vari risultati nella variabile creata in precedenza. In caso di errore ritorno null e l'errore
		if err := rows.Scan(&fabrizio.Codice, &inizio, &fine); err != nil {
			return nil, err
		}
		fabrizio.Inizio = time.Unix(inizio, 0)
		fabrizio.Fine = time.Unix(fine, 0)

		//Copio la variabile temporanea nell'ultima posizione dell'array
		prens = append(prens, fabrizio)
	}
	//Se c'è un errore, ritorna null e l'errore
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prens, nil
}

func AddImmagine(percorso, descrizione string, stanza uint32) (uint32, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return 0, err
	}

	//Preparo il database per la query
	stmt, err := db_Connection.Prepare(`INSERT INTO Immagine VALUES (null, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//Eseguo la query e ne salvo i risultati
	res, err := stmt.Exec(percorso, descrizione, stanza)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func AddPrenotazione(data_inizio time.Time, data_fine time.Time) (uint32, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return 0, err
	}

	//Preparo il database per la query
	stmt, err := db_Connection.Prepare(`INSERT INTO Prenotazione VALUES (null, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//Eseguo la query e ne salvo i risultati
	res, err := stmt.Exec(data_inizio.Unix(), data_fine.Unix())
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func AddStanza(nome, dotazioni string, metratura uint32) (uint32, error) {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return 0, err
	}

	//Preparo il database per la query
	stmt, err := db_Connection.Prepare(`INSERT INTO Stanza VALUES (null, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//Eseguo la query e ne salvo i risultati
	res, err := stmt.Exec(nome, dotazioni, metratura)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}


func SetNome(codice uint32, nome string) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `UPDATE Stanza
		  SET Nome = ?
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, nome, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func SetDotazioni(codice uint32, dotazioni string) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `UPDATE Stanza
		  SET Dotazioni = ?
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, dotazioni, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func SetPercorso(codice uint32, percorso string) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `UPDATE Immagine
		  SET Percorso = ?
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, percorso, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func SetDescrizione(codice uint32, descrizione string) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `UPDATE Immagine
		  SET Descrizione = ?
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, descrizione, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func SetStanza(codice, stanza uint32) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `UPDATE Immagine
		  SET Stanza = ?
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, stanza, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func RemoveImmagine(codice uint32) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `DELETE FROM Immagine
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func RemoveStanza(codice uint32) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `DELETE FROM Stanza
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}

func RemovePrenotazione(codice uint32) error {
	//Verifico se il server è ancora disponibile
	//Se c'è un errore, ritorna null e l'errore
	if err := db_Connection.Ping(); err != nil {
		return err
	}

	q := `DELETE FROM Prenotazione
		  WHERE Codice = ?`
	rows, err := db_Connection.Query(q, codice)
	//Se c'è un errore, ritorna null e l'errore
	if err != nil {
		return err
	}
	//Rows verrà chiuso una volta che tutte le funzioni normali saranno terminate oppure al prossimo return
	defer rows.Close()

	return nil
}
