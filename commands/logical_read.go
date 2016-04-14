package commands

import "github.com/spf13/cobra"

func (l *Logical) AddReadSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "Reads the value of the key at the given path.",
		Long:  "Reads the value of the key at the given path.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return l.Read(args)
		},
	}

	l.AddCommand(cmd)
}

func (l *Logical) Read(args []string) error {
	if err := l.CheckArgs(args); err != nil {
		return err
	}

	logical, err := l.Logical()
	if err != nil {
		return err
	}

	result, err := logical.Read(args[0])
	if err != nil {
		return err
	}

	return l.Output(result.Data["value"])
}
