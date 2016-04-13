package commands

import (
	"encoding/json"
	"fmt"
)

func (c *Cmd) Output(v interface{}) error {
	var err error
	var jsonRaw []byte

	jsonRaw, err = json.MarshalIndent(v, "", "  ")

	if err != nil {
		return err
	}

	fmt.Fprintf(c.Out, string(jsonRaw))

	return nil
}
