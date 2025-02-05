package add

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Used to add entity to the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AddCmd.AddCommand(addTruckCmd)
	AddCmd.AddCommand(addTrailerCmd)

}
