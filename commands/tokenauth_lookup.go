package commands

import (
	"log"

	"github.com/spf13/cobra"
)

// If you do c.AddCommand(), it will be a subcommand to tokenauth.
// If you do t.AddCommand(), it will be a root command.
// You need to set VAULT_TOKEN environment variables first
// before this can work. And export VAULT_ADDR=http://localhost:8200/ .
func (t *TokenAuth) AddLookupSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "lookup <token>",
		Short: "Display information about the specified token",
		Long:  "Display information about the specified token",
		RunE: func(cmd *cobra.Command, args []string) error {
			return t.Lookup(args)
		},
	}

	c.AddCommand(cmd)
}

func (t *TokenAuth) Lookup(args []string) error {
	if err := t.CheckArgs(args); err != nil {
		return err
	}

	// This comes from vault.go > TokenAuth()
	token, err := t.TokenAuth()
	if err != nil {
		return err
	}

	lookup, err := token.Lookup(args[0])
	if err != nil {
		log.Fatal(err)
	}
	return t.Output(lookup)
}
