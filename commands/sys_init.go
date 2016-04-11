package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddInitStatusSub(c *cobra.Command) {
	lookupCmd := &cobra.Command{
		Use:   "init-status",
		Short: "Return the initialization status of a Vault.",
		Long:  "Return the initialization status of a Vault.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.InitStatus(args)
		},
	}

	s.AddCommand(lookupCmd)
}

func (s *Sys) InitStatus(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.InitStatus()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return s.Output(result)
}
