/*
 * hash.go
 *
 * Pacchetto per la generazione e verifica dei codici dei libri.
 *
 * Copyright (c) 2022 Davide Vendramin <davidevendramin5@gmail.com>
 */

// Pacchetto per la generazione e verifica dei codici dei libri.
package hash

import (
	_ "crypto/rand"
	_ "crypto/sha256"
	_ "encoding/base64"
)

type ErrHash struct{}

func (e ErrHash) Error() string {
	return "Book hash didn't match"
}

/*func Verifica(pass string) (db.Libro, error) {
	pass_decoded, err := base64.RawURLEncoding.DecodeString(pass)
	if err != nil {
		return db.Libro{}, err
	}

	codice := uint32(0)
	for i := uint8(0); i < 4; i++ {
		codice += uint32(pass_decoded[i]) << (i * 8)
	}

	libro, err := db.GetLibro(codice)
	if err != nil {
		return libro, err
	}

	hash := sha256.Sum256(pass_decoded)
	if base64.StdEncoding.EncodeToString(hash[:]) != libro.Hashz {
		return libro, ErrHash{}
	}

	return libro, nil
}

func Genera(codice uint32) (string, error) {
	pass := make([]byte, 20)

	temp := codice
	for i := uint8(0); temp != 0; i++ {
		pass[i] = byte(temp & 0xFF)
		temp /= 256
	}

	if _, err := rand.Read(pass[4:]); err != nil {
		return "", err
	}

	pass_encoded := base64.RawURLEncoding.EncodeToString(pass)

	hash := sha256.Sum256(pass[:])
	db.SetHash(codice, base64.StdEncoding.EncodeToString(hash[:]))

	return pass_encoded, nil
}*/
