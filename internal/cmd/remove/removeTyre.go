/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreCmd represents the tyre command
var tyreCmd = &cobra.Command{
	Use:   "tyre",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyre called")
		tyreID := entries.ReadString("Please enter tyreID: ")
		err := entries.RemoveTyre(tyreID)
		if err != nil {
			fmt.Println("Error removing tyre:", err)
		}
	},
}

func init() {
	RemoveCmd.AddCommand(tyreCmd)

}
