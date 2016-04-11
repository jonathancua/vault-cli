package commands

import "github.com/spf13/cobra"

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
