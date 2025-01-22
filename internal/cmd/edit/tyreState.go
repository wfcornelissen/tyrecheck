/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package edit

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tyreStateCmd represents the tyreState command
var tyreStateCmd = &cobra.Command{
	Use:   "tyreState",
	Short: "Calls the tyre state entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre state entry function which asks for the tyre state.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyreState called")
	},
}

func init() {
	EditCmd.AddCommand(tyreStateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyreStateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyreStateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
