package interactive

import (
	"fmt"
	"github.com/spf13/cobra"
)

// InteractiveCmd allows user-input via interactive terminal prompts
var InteractiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive prompt",
	Long:  `Manage CA and Certificates using an interactive prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("interactive called")
	},
}

func init() {
	InteractiveCmd.AddCommand()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interactiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interactiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
