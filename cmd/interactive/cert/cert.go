package cert

import (
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"os"
	"runtime"
)

var certConfig = &tls.CertConfig{}
var caName string
var caPath = caPathByOS()

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
			log.Println("an error has occured: %w", err)
			return
		}
		selectUI := promptui.Select{
			Label: "Choose operation",
			Items: []string{"sign"},
		}
		_, res, err := selectUI.Run()
		if err != nil {
			log.Println("an error has occured: %w", err)
			return
		}
		switch res {
		case "sign":
			err = signUI()
		}
		if err != nil {
			log.Println("an error has occured: %w", err)
			return
		}
	},
}

func caPathByOS() (path string) {
	homeDir, _ := os.UserHomeDir()
	switch osEnv := runtime.GOOS; osEnv {
	case "windows":
		path = homeDir + "\\.simpleca"
	default:
		path = homeDir + "/.simpleca"
	}
	return
}
