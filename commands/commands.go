package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Cmd struct {
	root *cobra.Command

	Err io.Writer
	Out io.Writer

	vault *vault
}

func Init(name, version string) *Cmd {
	c := Cmd{
		Err:   os.Stderr,
		Out:   os.Stdout,
		vault: &vault{},
	}

	c.root = &cobra.Command{
		Use:   "vault-cli",
		Short: "Command line interface for Vault HTTP API",
		Long:  "Command line interface for Vault HTTP API",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c.root.Help()
			return nil
		},
	}

	c.root.PersistentFlags().StringVar(&c.vault.env, "env", "stage", "default environment")
	c.root.PersistentFlags().StringVar(&c.vault.address, "vault", "", "Vault address:port")
	c.root.PersistentFlags().BoolVar(&c.vault.sslSkipVerify, "ssl-verify", true, "Verify certificates when connecting via SSL")
	c.root.PersistentFlags().StringVar(&c.vault.sslClientCert, "ssl-client-cert", "", "Path to an SSL client certificate for authentication")
	c.root.PersistentFlags().StringVar(&c.vault.sslClientKey, "ssl-client-key", "", "Path to an SSL client certificate key for authentication")
	c.root.PersistentFlags().StringVar(&c.vault.sslCaCert, "ssl-ca-cert", "", "CA certificate file to validate the vault server")
	c.root.PersistentFlags().StringVar(&c.vault.sslCaPath, "ssl-ca-path", "", "Path to a CA certificate file")
	c.root.PersistentFlags().StringVar(&c.vault.token, "token", "", "The Vault token")

	c.initInfo()
	c.initLogical()
	c.initSys()
	c.initTokenAuth()

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information.",
		Long:  "Print version information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s %s\n", name, version)
			return nil
		},
	}
	c.root.AddCommand(versionCmd)

	return &c
}

func (c *Cmd) Execute() error {
	return c.root.Execute()
}

func (c *Cmd) Help() error {
	return c.root.Help()
}

func (c *Cmd) AddCommand(cmd *cobra.Command) {
	c.root.AddCommand(cmd)
}

type funcVar func(s string) error

func (f funcVar) Set(s string) error { return f(s) }
func (f funcVar) String() string     { return "" }
func (f funcVar) IsBoolFlag() bool   { return false }
func (f funcVar) Type() string       { return "funcVar" }
