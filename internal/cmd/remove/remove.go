package remove

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an entry from the database",
	Long:  `Remove a truck or trailer from the database based on fleet number input given by user`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
