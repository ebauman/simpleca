package ca

import (
	"github.com/ebauman/simpleca/file"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
)

var caPath = file.DefaultConfPath()

// Caprompt contains interactive prompt logic for CA management
var Caprompt = &cobra.Command{
	Use:   "ca",
	Short: "Interactive CA management",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Choose operation",
			Items: []string{"init"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			log.Println(err)
			return
		}

		switch res {
		case "init":
			err = initUI()
		}
		if err != nil {
			log.Println(err)
		}
	},
}
