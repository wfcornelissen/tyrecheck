package work

import (
	"github.com/spf13/cobra"
)

// workCmd represents the work command
var WorkCmd = &cobra.Command{
	Use:   "work",
	Short: "Used with subcommands to record work done on tyres",
	Long: `Used with subcommands to record work done on tyres
	
	Subcommands:
	create - create a new work entry
	read - read a work entry
	update - update a work entry
	delete - delete a work entry
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
