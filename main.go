package main

import (
	"github.com/spf13/cobra"
	"os"
	"tabley/cmd"
)

func main() {
	if err := tableyCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

// tableyCmd constructs and returns the cobra command for running tabley.
func tableyCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "tabley",
		Short: "Print music in tablature and other representations",
	}
	root.PersistentFlags().BoolP("basic", "b", false, "Print output in basic, sequential format.")

	scale := &cobra.Command{
		RunE:  cmd.ScaleRun,
		Use:   "scale <note>-<scale>",
		Short: "Prints the given scale",
		Long: `Scale takes a given scale (either preset or custom) and prints it in either basic or tablature format.

Available scales: 
- <note>-Major
- <note>-Minor[Harmonic/Natural]
- <note>-MajorPentatonic
- <note>-MinorPentatonic`,
	}

	root.AddCommand(scale)

	return root
}
