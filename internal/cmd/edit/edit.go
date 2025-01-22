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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
