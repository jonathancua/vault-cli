package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddListMountsSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "mounts",
		Short: "Lists all the mounted secret backends.",
		Long:  "Lists all the mounted secret backends.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.ListMounts(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) ListMounts(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.ListMounts()
	if err != nil {
		log.Fatal(err)
	}

	return s.Output(result)
}
