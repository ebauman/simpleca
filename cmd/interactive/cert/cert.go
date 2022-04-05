package cert

import (
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
)

var certConfig = &tls.CertConfig{}
var caName string
var caPath = file.ConfPathByOS()

var Certprompt = &cobra.Command{
	Use:   "cert",
	Short: "Interactive Certificate management",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		prompt := promptui.Prompt{
			Label:   "CA Name",
			Default: "default",
		}
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
