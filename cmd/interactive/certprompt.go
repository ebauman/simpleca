package interactive

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	Interactivecmd.AddCommand(Certprompt)
}

var Certprompt = &cobra.Command{
	Use:   "cert",
	Short: "Interactive Certificate management",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Is this interactive Certificate management prompt?",
			Items: []string{"yes", "no"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			fmt.Errorf("an error has occured: %w", err)
		}
		fmt.Println("Result:", res)
	},
}
