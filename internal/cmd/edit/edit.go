package edit

import (
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a tyre in the database",
	Long: `Starts function that asks for attributes of a tyre that is 
	already in the database.
	
	attributes are:
	Tyre Condition
	Tyre Location
	Tyre Position
	Tyre State
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
