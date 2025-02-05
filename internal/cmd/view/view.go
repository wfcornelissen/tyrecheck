package view

import (
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View entries",
	Long: `Used to view an entry's details. Can be
	one of the following:
	- tyre
	- truck
	- trailer`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
