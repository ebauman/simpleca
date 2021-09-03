package ca

import (
	"fmt"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var includePrivkey bool

func init() {
	CAcmd.AddCommand(viewCmd)

	viewCmd.Flags().BoolVar(&includePrivkey, "include-privkey", false,
		"include private key when viewing certificate")
}

var viewCmd = &cobra.Command {
	Use: "view",
	Short: "view ca certificate",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var caname string
		if len(args) > 0 {
			caname = args[0]
		} else {
			caname = "default"
		}

		certPath := fmt.Sprintf("%s/%s/%s", capath, caname, tls.CACertFileName)
		keyPath := fmt.Sprintf("%s/%s/%s", capath, caname, tls.CAKeyFileName)

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