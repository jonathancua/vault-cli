package commands

import "github.com/spf13/cobra"

func (s *Sys) AddLeaderSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "leader",
		Short: "Returns the high availability status and current leader instance of Vault.",
		Long:  "Returns the high availability status and current leader instance of Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.Leader(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) Leader(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		return err
	}

	result, err := sys.Leader()
	if err != nil {
		return err
	}

	return s.Output(result)
}
