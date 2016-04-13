package commands

import "github.com/spf13/cobra"

func (s *Sys) AddGetPolicySub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "get-policy",
		Short: "Retrieve the rules for the named policy.",
		Long:  "Retrieve the rules for the named policy.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.GetPolicy(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) GetPolicy(args []string) error {
	if err := s.CheckArgs(args); err != nil {
		return err
	}

	sys, err := s.Sys()
	if err != nil {
		return err
	}

	result, err := sys.GetPolicy(args[0])
	if err != nil {
		return err
	}

	return s.Output(result)
}
