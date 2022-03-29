/*
 * ldap.go
 *
 * Autenticazione su un server LDAP.
 *
 * Copyright (c) 2022 Davide Vendramin <davidevendramin5@gmail.com>
 */

// Package per l'autenticazione degli utenti.
package auth

import (
	"errors"
	"fmt"
	"log"

	"github.com/pentecoste/webcanarie/internal/config"
	ldap "github.com/go-ldap/ldap/v3"
)

// Errore di autenticazione
type AuthenticationError struct {
	user string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication error for user %s.", e)
}

// Controlla le credenziali sul server LDAP
func checkCredentials(username string, password string) (UserInfo, error) {

	// Ottiene la configurazione
	host := config.Config.LDAP.URI
	port := config.Config.LDAP.Porta
	baseDN := config.Config.LDAP.BaseDN
	bindUserDN := "cn=" + config.Config.LDAP.Username + "," + baseDN
	bindPassword := config.Config.LDAP.Password

	// Connessione al server LDAP
	l, err := ldap.DialURL("ldap://" + host + ":" + port)
	if err != nil {
		log.Println("auth: ", err.Error())
		return dummyUserInfo, &AuthenticationError{username}
	}
	defer l.Close()

	// Per prima cosa effettuo l'accesso con un utente admin
	err = l.Bind(bindUserDN, bindPassword)
	if err != nil {
		log.Println("auth: ", err.Error())
		return dummyUserInfo, &AuthenticationError{username}
	}

	// Cerco l'username richiesto
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", username),
		[]string{"dn", "cn", "ou"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Println("auth: ", err.Error())
		return dummyUserInfo, &AuthenticationError{username}
	}

	// Verifico il numero di utenti corrispondenti e ottendo il DN dell'utente
	if len(sr.Entries) != 1 {
		return dummyUserInfo, &AuthenticationError{username}
	}

	userDN := sr.Entries[0].DN
	fullName := sr.Entries[0].GetAttributeValue("cn")
	group := sr.Entries[0].GetAttributeValue("ou")

	// Verifica la password
	err = l.Bind(userDN, password)
	if err != nil {
		return dummyUserInfo, errors.New("Password errata!")
	}

	return UserInfo{username, fullName, group}, nil
}
