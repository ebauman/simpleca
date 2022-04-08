package cert

import (
	"crypto/x509"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vltraheaven/simpleca/file"
	"github.com/vltraheaven/simpleca/table"
	"github.com/vltraheaven/simpleca/tls"
	"os"
)

func init() {
	Certcmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list certificates",
	RunE: func(cmd *cobra.Command, args []string) error {
		fullPath := fmt.Sprintf("%s/%s", certConfig.Path, caName)

		dirs, err := file.ListDirectories(fullPath)
		if err != nil {
			return err
		}

		// for each directory (potential cert), try and load it
		validCerts := map[string]*x509.Certificate{}

		for _, d := range dirs {
			var certPath = fmt.Sprintf("%s/%s/%s", fullPath, d, tls.CertFileName)
			var keyPath = fmt.Sprintf("%s/%s/%s", fullPath, d, tls.KeyFileName)

			_, cert, err := tls.LoadCert(certPath, keyPath)
			if err != nil {
				return err
			}

			validCerts[d] = cert
		}

		return table.WriteCAList(os.Stdout, validCerts)
	},
}
