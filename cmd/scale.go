package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// ScaleRun starts the core logic of the Scale subcommand.
func ScaleRun(cmd *cobra.Command, args []string) error {
	for _, x := range args {
		fmt.Println(x)
	}
	_, err := cmd.Flags().GetBool("basic")
	if err != nil {
		return errors.Join(err, errors.New("failed to get basic flag"))
	}

	//musicProcessor := processor.NewProcessorFromArg(basic)

	return nil
}
