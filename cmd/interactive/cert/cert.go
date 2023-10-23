package cert

import (
	"fmt"
	"github.com/ebauman/simpleca/cmd/interactive/ca"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/parse"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
)

var caName string
var caPath = file.DefaultConfPath()

// Certprompt contains interactive prompt logic for certificate management
var Certprompt = &cobra.Command{
	Use:   "cert",
	Short: "Interactive Certificate management",
	Run: func(cmd *cobra.Command, args []string) {
		promptPath := promptui.Prompt{
			Label:    "Enter path to CA storage directory",
			Validate: parse.ValidatePath,
			Default:  caPath,
		}
		p, err := promptPath.Run()
		if err != nil {
			log.Println(err)
			return
		}
		if p != "" || p != caPath {
			caPath = p
		}

		cas, err := ca.ListCAs(caPath)
		if err != nil {
			log.Println(err)
			return
		}

		selectUI := promptui.Select{
			Label: "Choose CA",
			Items: cas,
		}
		_, caName, err = selectUI.Run()
		if err != nil {
			log.Println(err)
			return
		}
		selectUI = promptui.Select{
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
			fmt.Println(err)
		}
		return
	},
}
