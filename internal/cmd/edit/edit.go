package edit

import (
	"fmt"

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
		fmt.Println("edit called")
	},
}

func init() {

}
