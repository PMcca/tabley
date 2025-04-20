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
	root.PersistentFlags().BoolP(
		cmd.FlagBasic,
		"b",
		false,
		"Print output in basic, sequential format.",
	)

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

	tab := &cobra.Command{
		RunE:  cmd.TabRun,
		Use:   "tab <tablature>",
		Short: "Convert a given guitar tab's fret numbers to their corresponding notes",
		Long: `Tab takes a guitar tab and prints it with each fret number converted to their corresponding note.

The tuning can be defined using the -t flag. This requires each note that each string is tuned in. For example, -t DADGBE.
If -t is omitted, standard tuning (EADGBE) will be used.
`,
	}
	tab.LocalFlags().StringP(
		cmd.FlagTuning,
		"t",
		"EADGBE",
		"Tuning (in notes) for this tab.",
	)

	root.AddCommand(scale, tab)

	return root
}

//package main
//
//import (
//	"fmt"
//)
//
//func transformString(row string) string {
//	runes := []rune(row)
//	result := make([]rune, 0, len(runes)*2)
//
//	for _, char := range runes {
//		switch char {
//		case '$':
//			result = append(result, '9')
//		case '£':
//			result = append(result, '1', '0') // Replace £ with "10" using two characters
//		default:
//			result = append(result, char)
//		}
//	}
//
//	return string(result)
//}
//
//func main() {
//	row1 := "---$----$--$---$-"
//	row2 := "-----$--$--£-£-£-"
//
//	fmt.Println(transformString(row1))
//	fmt.Println(transformString(row2))
//}
