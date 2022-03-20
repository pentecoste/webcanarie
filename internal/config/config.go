/*
 * config.go
 *
 * File per il caricamento e gestione della configurazione
 *
 * Copyright (c) 2022 Davide Vendramin <davidevendramin5@gmail.com>
 */

// File per il caricamento e gestione della configurazione
package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type config struct {
	Generale       generale       `toml:"Generale"`
	Autenticazione autenticazione `toml:"Autenticazione"`
	LDAP           ldap           `toml:"LDAP"`
	SQL            sql            `toml:"SQL"`
}

type generale struct {
	FQDN            string `toml:"fqdn_sito"`
	Porta           string `toml:"porta_http"`
	AdminUser       string `toml:"utente_admin"`
	LunghezzaPagina uint16 `toml:"lunghezza_pagina"`
}

type autenticazione struct {
	JWTSecret     string `toml:"chiave_firma"`
	SSO           bool   `toml:"usa_sso"`
	SSOURL        string `toml:"sso_url"`
	SecureCookies bool   `toml:"cookie_sicuri"`
	DummyAuth     bool   `toml:"dummy_auth"`
}

type ldap struct {
	URI      string `toml:"uri"`
	Porta    string `toml:"porta"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	BaseDN   string `toml:"base_dn"`
}

type sql struct {
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Indirizzo string `toml:"indirizzo"`
	Database  string `toml:"database"`
}

var Config config

func LoadConfig(path string) error {
	absPath, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(absPath, &Config); err != nil {
		return err
	}

	return nil
}
