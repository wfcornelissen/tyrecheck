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
		TyreID := entries.ReadString("Enter the tyre ID: ")
		entries.EditPosition(TyreID)
	},
}

func init() {
	EditCmd.AddCommand(PositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyrePositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyrePositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
