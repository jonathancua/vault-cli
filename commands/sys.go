package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type Sys struct {
	*Cmd
}

func (root *Cmd) initSys() {
	s := Sys{Cmd: root}

	sysCmd := &cobra.Command{
		Use:   "sys",
		Short: "Vault /sys endpoint interface",
		Long:  "Vault /sys endpoint interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.sysInitStatus(args)
		},
	}

	s.AddCommand(sysCmd)
}

func (s *Sys) sysInitStatus(args []string) error {
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
