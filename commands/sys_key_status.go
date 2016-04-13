package commands

import "github.com/spf13/cobra"

func (s *Sys) AddKeyStatusSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "key-status",
		Short: "Returns information about the current encryption key used by Vault.",
		Long:  "Returns information about the current encryption key used by Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.KeyStatus(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) KeyStatus(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		return err
	}

	result, err := sys.KeyStatus()
	if err != nil {
		return err
	}

	return s.Output(result)
}
