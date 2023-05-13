package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"tabley/processor"
)

// TabRun starts the core logic of the Tab subcommand.
func TabRun(cmd *cobra.Command, args []string) error {
	tuning, err := cmd.LocalFlags().GetString(FlagTuning)
	if err != nil {
		return errors.Wrap(err, "failed to get tuning flag")
	}

	basic, err := cmd.Flags().GetBool(FlagBasic)
	if err != nil {
		return errors.Wrap(err, "failed to get basic flag")
	}

	musicProcessor := processor.NewProcessorFromArg(basic)
	tab, err := musicProcessor.ConvertTab(args, tuning)
	if err != nil {
		return errors.Wrap(err, "failed to process tab")
	}

	fmt.Println(tab)
	return nil
}
