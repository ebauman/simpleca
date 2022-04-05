package interactive

import (
	"fmt"
	"github.com/ebauman/simpleca/cmd/interactive/ca"
	"github.com/ebauman/simpleca/cmd/interactive/cert"
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
			ca.Caprompt.Run(cmd, args)
		case "cert":
			cert.Certprompt.Run(cmd, args)
		}
	},
}

func init() {
	Interactivecmd.AddCommand(cert.Certprompt)
	Interactivecmd.AddCommand(ca.Caprompt)
}
