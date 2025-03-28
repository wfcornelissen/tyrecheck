package work

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		// Call workFuncs from here
		err := entries.RemoveTyreWork(args[0])
		if err != nil {
			fmt.Println("Error removing tyre work:", err)
		}
		fmt.Println("Tyre work removed")
	},
}

func init() {
	WorkCmd.AddCommand(removeCmd)
}
