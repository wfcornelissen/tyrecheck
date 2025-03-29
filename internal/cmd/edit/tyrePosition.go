package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyrePositionCmd represents the tyrePosition command
var PositionCmd = &cobra.Command{
	Use:   "position",
	Short: "Calls the tyre position entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre position entry function which asks for the tyre position.`,
	Run: func(cmd *cobra.Command, args []string) {
		entries.EditPosition()
	},
}

func init() {
	EditCmd.AddCommand(PositionCmd)

}
