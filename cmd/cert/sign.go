package cert

import (
	"github.com/ebauman/simpleca/cli"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
)

func init() {
	Certcmd.AddCommand(signCmd)

	cli.LoadCertFlags(certConfig, signCmd)
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "sign new certificates from a ca",
	Args:  cobra.MinimumNArgs(1), // vanity name of the cert
	RunE: func(cmd *cobra.Command, args []string) error {
		certConfig.Name = args[0]
		return tls.SignCert(certConfig, caName)
	},
}
