package ca

import (
	"fmt"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var privKey bool

func init() {
	CAcmd.AddCommand(viewCmd)

	viewCmd.Flags().BoolVar(&privKey, "key", false,
		"view private key instead of certificate")
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view ca certificate",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var caname string
		if len(args) > 0 {
			caname = args[0]
		} else {
			caname = "default"
		}

		certPath := fmt.Sprintf("%s/%s/%s", capath, caname, tls.CACertFileName)
		keyPath := fmt.Sprintf("%s/%s/%s", capath, caname, tls.CAKeyFileName)

		if privKey {
			keyBytes, err := os.ReadFile(keyPath)
			if err != nil {
				return err
			}

			fmt.Print(string(keyBytes))

			return nil
		}

		certBytes, err := os.ReadFile(certPath)
		if err != nil {
			return err
		}

		fmt.Print(string(certBytes))

		return nil
	},
}
