package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreConditionCmd represents the tyreCondition command
var ConditionCmd = &cobra.Command{
	Use:   "condition",
	Short: "Calls the tyre condition entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre condition entry function which asks for the tyre condition.`,
	Run: func(cmd *cobra.Command, args []string) {
		entries.EditCondition()
	},
}

func init() {
	EditCmd.AddCommand(ConditionCmd)

}
