package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func (s *Sys) AddListAuditSub(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "audits",
		Short: "Lists the mounted audit backends.",
		Long:  "Lists the mounted audit backends.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.ListAudit(args)
		},
	}

	s.AddCommand(cmd)
}

func (s *Sys) ListAudit(args []string) error {
	sys, err := s.Sys()
	if err != nil {
		log.Fatal(err)
	}

	result, err := sys.ListAudit()
	if err != nil {
		log.Fatal(err)
	}

	return s.Output(result)
}
