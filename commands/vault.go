package commands

import (
	"crypto/tls"

	vaultapi "github.com/hashicorp/vault/api"
)

type vault struct {
	configFile    string
	env           string
	address       string
	sslCaCert     string
	sslCaPath     string
	sslClientCert string
	sslClientKey  string
	sslVerify     bool
	token         string
	tlsConfig     *tls.Config
}

// Because Client type has Sys() method which return Sys type
func (c *Cmd) Sys() (*vaultapi.Sys, error) {
	vault, err := c.Client()
	if err != nil {
		return nil, err
	}

	return vault.Sys(), nil
}

// type Client has method Auth() but Auth() has method Token()
// and it returns *TokenAuth.
// See how it relates to the below return value and return type
func (c *Cmd) TokenAuth() (*vaultapi.TokenAuth, error) {
	vault, err := c.Client()
	if err != nil {
		return nil, err
	}

	return vault.Auth().Token(), nil
}

// Note how I added a method for the Cmd type,
// and the method is related to vault
func (c *Cmd) Client() (*vaultapi.Client, error) {
	config := vaultapi.DefaultConfig()
	vsl := c.vault
	vsl.tlsConfig = new(tls.Config)

	if vsl.address != "" {
		config.Address = c.vault.address
	}

	client, err := vaultapi.NewClient(config)

	if err != nil {
		return nil, err
	}

	if vsl.token != "" {
		client.SetToken(vsl.token)
	}

	return client, nil
}

func NewVault() *vault {
	return &vault{}
}
