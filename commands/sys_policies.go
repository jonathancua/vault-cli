package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddListPoliciesSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "policies",
		Short: "Lists all the available policies.",
		Long:  "Lists all the available policies.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.ListPolicies(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) ListPolicies(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.ListPolicies()
	if err != nil {
		log.Fatal(err)
	}

	return s.Output(result)
}
