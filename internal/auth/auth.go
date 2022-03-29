/*
 * auth.go
 *
 * Funzione per autenticare un utente.
 *
 * Copyright (c) 2022 Davide Vendramin <davidevendramin5@gmail.com>
 */

package auth

import "github.com/pentecoste/webcanarie/internal/config"

// Informazioni sull'utente
type UserInfo struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Group    string `json:"group"`
}

// Verifica le credenziali e restituisce il token.
func AuthenticateUser(username, password string) ([]byte, error) {

	var (
		err      error
		userInfo UserInfo
	)

	// Controlla le credenziali
	if !config.Config.Autenticazione.DummyAuth {
		userInfo, err = checkCredentials(username, password)
	} else {
		err = nil
		userInfo = UserInfo{username, "1337 h4x0r", "h4x0rz"}
	}

	if err == nil {
		// Genera il token
		token, err := getToken(userInfo)

		if err != nil {
			return nil, err
		}

		return token, nil

	} else {
		return nil, err
	}
}
