package commands

import "github.com/spf13/cobra"

func (l *Logical) AddListSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List keys at the given path.",
		Long:  "List keys at the given path.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return l.List(args)
		},
	}

	l.AddCommand(cmd)
}

func (l *Logical) List(args []string) error {
	if err := l.CheckArgs(args); err != nil {
		return err
	}

	logical, err := l.Logical()
	if err != nil {
		return err
	}

	result, err := logical.List(args[0])
	if err != nil {
		return err
	}

	return l.Output(result.Data["keys"])
}
