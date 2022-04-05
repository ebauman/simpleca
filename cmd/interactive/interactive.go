package interactive

import (
	"github.com/ebauman/simpleca/cmd/interactive/ca"
	"github.com/ebauman/simpleca/cmd/interactive/cert"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
)

// Interactivecmd allows user-input via interactive terminal prompts
var Interactivecmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive ca/cert management",
	Long:  `Manage CA and Certificates using an interactive prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Choose management prompt",
			Items: []string{"ca", "cert"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			log.Println("an error has occured: %w", err)
			return
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
