package commands

import (
	"crypto/tls"
	"fmt"
	"os"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

type vault struct {
	configFile    string
	env           string
	address       string
	sslCaCert     string
	sslCaPath     string
	sslClientCert string
	sslClientKey  string
	sslSkipVerify bool
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
	config.ReadEnvironment()
	configFile := c.GetConfig()

	vsl := c.vault
	vsl.tlsConfig = new(tls.Config)

	// hierarchy to check for the vault address:
	//  file, environment variable, flag
	config.Address = configFile.vaultAddr
	if os.Getenv("VAULT_ADDR") != "" {
		fmt.Println("Found VAULT_ADDR variable, using it")
		config.Address = os.Getenv("VAULT_ADDR")
	}
	if vsl.address != "" {
		config.Address = c.vault.address
	}

	client, err := vaultapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	clientToken := configFile.vaultToken
	if os.Getenv("VAULT_TOKEN") != "" {
		fmt.Println("Found VAULT_TOKEN  variable, using it")
		clientToken = os.Getenv("VAULT_TOKEN")
	}
	if vsl.token != "" {
		clientToken = vsl.token
	}
	client.SetToken(clientToken)

	if configFile.vaultCaCert != "" {
		vaultapi.LoadCACert(configFile.vaultCaCert)
	}

	return client, nil
}

type ConfigFromFile struct {
	vaultAddr   string
	vaultToken  string
	vaultCaCert string
}

func (c *Cmd) GetConfig() *ConfigFromFile {
	config := &ConfigFromFile{}
	viper.SetConfigName(".vault-cli")
	viper.AddConfigPath("$HOME")
	viper.ReadInConfig()

	vaultAddrStr := fmt.Sprintf("%s.vault_addr", c.vault.env)
	vaultTokenStr := fmt.Sprintf("%s.vault_token", c.vault.env)
	config.vaultAddr = viper.GetString(vaultAddrStr)
	config.vaultToken = viper.GetString(vaultTokenStr)
	config.vaultCaCert = viper.GetString("vault_cacert")

	return config
}

func NewVault() *vault {
	return &vault{}
}
