package commands

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

type Info struct {
	*Cmd
}

func (root *Cmd) initInfo() {
	i := Info{Cmd: root}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Give info about the config to be used.",
		Long:  "Give info about the config to be used.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.infoResult(args)
		},
	}

	i.AddCommand(cmd)
}

func (i *Info) infoResult(args []string) error {
	info, err := i.Client()
	if err != nil {
		return err
	}

	spew.Dump(info)

	return nil
}
