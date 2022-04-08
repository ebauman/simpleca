package cert

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/vltraheaven/simpleca/file"
	"log"
)

var caName string
var caPath = file.DefaultConfPath()

// Certprompt contains interactive prompt logic for certificate management
var Certprompt = &cobra.Command{
	Use:   "cert",
	Short: "Interactive Certificate management",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label:   "CA Name",
			Default: "default",
		}
		var err error
		caName, err = prompt.Run()
		if err != nil {
			log.Println(err)
			return
		}
		selectUI := promptui.Select{
			Label: "Choose operation",
			Items: []string{"sign"},
		}
		_, res, err := selectUI.Run()
		if err != nil {
			log.Println(err)
			return
		}
		switch res {
		case "sign":
			err = signUI()
		}
		if err != nil {
			log.Println(err)
			return
		}
	},
}
