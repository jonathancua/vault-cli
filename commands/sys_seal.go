package commands

import "github.com/spf13/cobra"

func (s *Sys) AddSealStatusSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "seal-status",
		Short: "Returns the seal status of the Vault.",
		Long:  "Returns the seal status of the Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.SealStatus(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) SealStatus(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		return err
	}

	result, err := sys.SealStatus()
	if err != nil {
		return err
	}

	return s.Output(result)
}
