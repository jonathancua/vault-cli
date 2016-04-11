package commands

import "github.com/spf13/cobra"

type Sys struct {
	*Cmd
}

func (root *Cmd) initSys() {
	s := Sys{Cmd: root}

	sysCmd := &cobra.Command{
		Use:   "sys",
		Short: "Vault /sys endpoint interface",
		Long:  "Vault /sys endpoint interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	s.AddInitStatusSub(sysCmd)
}
