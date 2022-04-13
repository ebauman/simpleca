package cert

import (
	"fmt"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var privKey bool

func init() {
	Certcmd.AddCommand(viewCmd)

	viewCmd.Flags().BoolVar(&privKey, "key", false,
		"view private key instead of certificate")
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view certificate",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// first (and only?) argument is the name of the cert to view
		fullPath := tls.FullCertPath(certConfig, caName)

		certPath := fmt.Sprintf("%s/%s/%s", fullPath, args[0], tls.CertFileName)
		keyPath := fmt.Sprintf("%s/%s/%s", fullPath, args[0], tls.KeyFileName)

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
