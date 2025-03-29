package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreLocationCmd represents the tyreLocation command
var LocationCmd = &cobra.Command{
	Use:   "location",
	Short: "Calls the tyre location entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre location entry function which asks for the tyre location.`,
	Run: func(cmd *cobra.Command, args []string) {
		entries.EditLocation()
	},
}

func init() {
	EditCmd.AddCommand(LocationCmd)

}
