package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type TokenAuth struct {
	*Cmd
}

func (root *Cmd) initTokenAuth() {
	t := TokenAuth{Cmd: root}

	tokenCmd := &cobra.Command{
		Use:   "token-auth",
		Short: "Vault /tokenauth endpoint interface",
		Long:  "Vault /tokenauth endpoint interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	t.AddLookupSub(tokenCmd)
	t.AddCommand(tokenCmd)
}

func (t *TokenAuth) CheckArgs(args []string) error {
	switch {
	case len(args) == 0:
		return fmt.Errorf("Token must be specified")
	case len(args) > 1:
		return fmt.Errorf("Only one token allowed")
	}

	return nil
}
