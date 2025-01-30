package view

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tyreCmd represents the tyre command
var tyreCmd = &cobra.Command{
	Use:   "tyre",
	Short: "View tyre details",
	Long:  `Displays details of tyre requested according to tyreID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyre called")
	},
}

func init() {
	ViewCmd.AddCommand(tyreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
