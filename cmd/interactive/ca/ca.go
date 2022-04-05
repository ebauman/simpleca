package ca

import (
	"fmt"
	"github.com/ebauman/simpleca/cmd/interactive"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	interactive.Interactivecmd.AddCommand(Caprompt)
}

var Caprompt = &cobra.Command{
	Use:   "ca",
	Short: "Interactive CA management",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Is this interactive CA management prompt?",
			Items: []string{"yes", "no"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			fmt.Errorf("an error has occured: %w", err)
		}
		fmt.Println("Result:", res)
	},
}
