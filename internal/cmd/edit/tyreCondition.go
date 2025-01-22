/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreConditionCmd represents the tyreCondition command
var tyreConditionCmd = &cobra.Command{
	Use:   "tyreCondition",
	Short: "Calls the tyre condition entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre condition entry function which asks for the tyre condition.`,
	Run: func(cmd *cobra.Command, args []string) {
		TyreID := entries.ReadString("Enter the tyre ID: ")
		entries.EditCondition(TyreID)
	},
}

func init() {
	EditCmd.AddCommand(tyreConditionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyreConditionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyreConditionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
