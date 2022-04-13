package ca

import (
	"crypto/x509"
	"fmt"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/table"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	CAcmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list certificate authorities",
	RunE: func(cmd *cobra.Command, args []string) error {
		// get a list of directories inside the specified path
		dirs, err := file.ListDirectories(capath)
		if err != nil {
			return err
		}

		// for each directory (potential CA), try and load the CA
		validCAs := map[string]*x509.Certificate{}

		for _, d := range dirs {
			var certPath = fmt.Sprintf("%s/%s/%s", capath, d, tls.CACertFileName)
			var keyPath = fmt.Sprintf("%s/%s/%s", capath, d, tls.CAKeyFileName)

			_, cert, err := tls.LoadCA(certPath, keyPath)
			if err != nil {
				return err
			}

			validCAs[d] = cert
		}

		return table.WriteCAList(os.Stdout, validCAs)
	},
}
