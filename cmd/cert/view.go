package cert

import (
	"fmt"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var includePrivkey bool

func init() {
	Certcmd.AddCommand(viewCmd)

	viewCmd.Flags().BoolVar(&includePrivkey, "include-privkey", false,
		"include private key when viewing certificate")
}

var viewCmd = &cobra.Command{
	Use: "view",
	Short: "view certificate",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// first (and only?) argument is the name of the cert to view
		fullPath := tls.FullCertPath(certConfig, caName)

		certPath := fmt.Sprintf("%s/%s/%s", fullPath, args[0], tls.CertFileName)
		keyPath := fmt.Sprintf("%s/%s/%s", fullPath, args[0], tls.KeyFileName)

		certBytes, err := os.ReadFile(certPath)
		if err != nil {
			return err
		}

		if includePrivkey {
			keyBytes, err := os.ReadFile(keyPath)
			if err != nil {
				return err
			}

			fmt.Println(string(keyBytes))
		}

		fmt.Println(string(certBytes))

		return nil
	},
}