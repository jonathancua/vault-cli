package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Sys struct {
	*Cmd
}

func (root *Cmd) initSys() {
	s := Sys{Cmd: root}

	cmd := &cobra.Command{
		Use:   "sys",
		Short: "Vault /sys endpoint interface",
		Long:  "Vault /sys endpoint interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	s.AddInitStatusSub(cmd)
	s.AddGetPolicySub(cmd)
	s.AddKeyStatusSub(cmd)
	s.AddLeaderSub(cmd)
	s.AddListAuthSub(cmd)
	s.AddListAuditSub(cmd)
	s.AddListMountsSub(cmd)
	s.AddListPoliciesSub(cmd)
	s.AddSealStatusSub(cmd)
}

func (s *Sys) CheckArgs(args []string) error {
	switch {
	case len(args) == 0:
		return fmt.Errorf("Policy must be specified")
	case len(args) > 1:
		return fmt.Errorf("Only one token allowed")
	}

	return nil
}
