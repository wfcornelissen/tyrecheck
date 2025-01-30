package view

import (
	"fmt"

	"github.com/spf13/cobra"
)

// truckCmd represents the truck command
var truckCmd = &cobra.Command{
	Use:   "truck",
	Short: "View truck details",
	Long:  `Displays details of truck requested according to fleet number.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("truck called")
	},
}

func init() {
	ViewCmd.AddCommand(truckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// truckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// truckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
