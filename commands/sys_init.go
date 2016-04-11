package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddInitStatusSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "init-status",
		Short: "Return the initialization status of the Vault.",
		Long:  "Return the initialization status of the Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.InitStatus(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) InitStatus(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.InitStatus()
	if err != nil {
		log.Fatal(err)
	}

	return s.Output(result)
}
