package interactive

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Interactivecmd allows user-input via interactive terminal prompts
var Interactivecmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive ca/cert management",
	Long:  `Manage CA and Certificates using an interactive prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Choose management target",
			Items: []string{"ca", "cert"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			fmt.Errorf("an error has occured: %w", err)
		}
		switch res {
		case "ca":
			Caprompt.Run(cmd, args)
		case "cert":
			Certprompt.Run(cmd, args)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interactiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interactiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
