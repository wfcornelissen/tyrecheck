/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package edit

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tyrePositionCmd represents the tyrePosition command
var tyrePositionCmd = &cobra.Command{
	Use:   "tyrePosition",
	Short: "Calls the tyre position entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre position entry function which asks for the tyre position.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyrePosition called")
	},
}

func init() {
	EditCmd.AddCommand(tyrePositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyrePositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyrePositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
