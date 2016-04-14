package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Logical struct {
	*Cmd
}

func (root *Cmd) initLogical() {
	l := Logical{Cmd: root}

	cmd := &cobra.Command{
		Use:   "logical",
		Short: "Vault /logical endpoint interface",
		Long:  "Vault /logical endpoint interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	l.AddReadSub(cmd)
}

func (l *Logical) CheckArgs(args []string) error {
	switch {
	case len(args) == 0:
		return fmt.Errorf("Read path must be specified")
	case len(args) > 1:
		return fmt.Errorf("Only one path can be specified")
	}

	return nil
}
