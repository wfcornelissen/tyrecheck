package add

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreCmd represents the tyre command
var tyreCmd = &cobra.Command{
	Use:   "tyre",
	Short: "Calls the tyre entry function",
	Long: `Is used as subcommand for add to call
	the tyre entry function.
	
	It will ask for the following information:
	- Tyre size
	- Tyre type
	- Tyre brand
	- Tyre position
	- Tyre condition
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyre called")
		entries.AddTyre()
	},
}

func init() {
	AddCmd.AddCommand(tyreCmd)

}
