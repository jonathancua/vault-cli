package commands

import (
	"encoding/json"
	"fmt"
)

func (c *Cmd) Output(v interface{}) error {
	return c.OutputJSON(v, true)
}

func (c *Cmd) OutputJSON(v interface{}, prettyFlag bool) error {
	var err error
	var jsonRaw []byte

	if prettyFlag {
		jsonRaw, err = json.MarshalIndent(v, "", "  ")
	} else {
		jsonRaw, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	fmt.Fprintf(c.Out, string(jsonRaw))

	return nil
}
