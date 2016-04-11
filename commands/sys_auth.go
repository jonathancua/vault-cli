package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddListAuthSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "auths",
		Short: "Lists all the enabled auth backends.",
		Long:  "Lists all the enabled auth backends.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.ListAuth(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) ListAuth(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.ListAuth()
	if err != nil {
		log.Fatal(err)
	}

	return s.Output(result)
}
