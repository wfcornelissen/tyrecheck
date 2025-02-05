/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package link

import (
	"fmt"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Used to link entities together",
	Long: `Used to link entities together.
	-Truck and Trailer
	-Truck and Tyre
	-Trailer and Tyre`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("link called")
	},
}

func init() {

}
