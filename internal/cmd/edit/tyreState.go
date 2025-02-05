package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreStateCmd represents the tyreState command
var StateCmd = &cobra.Command{
	Use:   "state",
	Short: "Calls the tyre state entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre state entry function which asks for the tyre state.`,
	Run: func(cmd *cobra.Command, args []string) {
		TyreID := entries.ReadString("Enter the tyre ID: ")
		entries.EditState(TyreID)
	},
}

func init() {
	EditCmd.AddCommand(StateCmd)
}
