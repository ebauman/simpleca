package ca

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var deleteConfirm bool

func init() {
	CAcmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVar(&deleteConfirm, "confirm", false,
		"auto confirmation of ca deletion")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete certificate authority",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		caPath := fmt.Sprintf("%s/%s", capath, args[0])

		reader := bufio.NewReader(os.Stdin)

		if !deleteConfirm {
			fmt.Print("Are you sure you want to delete this CA? [y/N]: ")
			answer, _ := reader.ReadString('\n')
			answer = strings.Replace(answer, "\n", "", -1)

			if answer != "y" && answer != "Y" && answer != "yes" {
				fmt.Println("didn't receive confirmation, doing nothing")
				return nil
			}
		}

		// confirmed either by flag or input, so delete

		return os.RemoveAll(caPath)
	},
}
