package cert

import (
	"bufio"
	"fmt"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var deleteConfirm bool

func init() {
	Certcmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVar(&deleteConfirm, "confirm", false,
		"auto confirmation of certificate deletion")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete certificate",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fullPath := tls.FullCertPath(certConfig, caName)

		certPath := fmt.Sprintf("%s/%s", fullPath, args[0])

		reader := bufio.NewReader(os.Stdin)

		if !deleteConfirm {
			fmt.Print("Are you sure you want to delete this certificate? [y/N]: ")
			answer, _ := reader.ReadString('\n')
			answer = strings.Replace(answer, "\n", "", -1)

			if answer != "y" && answer != "Y" && answer != "yes" {
				fmt.Println("didn't receive confirmation, doing nothing")
				return nil
			}
		}

		// confirmed either by flag or by input, so delete

		return os.RemoveAll(certPath)
	},
}
