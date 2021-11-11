package cli

import (
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"net"
)

func LoadCertFlags(certConfig *tls.CertConfig, cmd *cobra.Command) {
	cmd.Flags().StringVar(&certConfig.Passphrase, "passphrase", "changeme",
		"passphrase for generated certificate")
	cmd.Flags().StringVar(&certConfig.Country, "country", "AA", "2-letter country code")
	cmd.Flags().StringVar(&certConfig.State, "state", "Relaxation", "state or province name")
	cmd.Flags().StringVar(&certConfig.Locality, "locality", "", "locality")
	cmd.Flags().StringVar(&certConfig.Organization, "organization", "SimpleCA Ltd.",
		"organization name")
	cmd.Flags().StringVar(&certConfig.OrganizationalUnit, "organizational-unit", "SimpleCA Security",
		"organizational unit")
	cmd.Flags().IPSliceVar(&certConfig.IPAddresses, "ip", []net.IP{},
		"ip address subject alternative name")
	cmd.Flags().StringSliceVar(&certConfig.DNSNames, "dns", []string{}, "dns subject alternative name")
	cmd.Flags().StringSliceVar(&certConfig.EmailAddresses, "email", []string{},
		"email subject alternative name")
	cmd.Flags().StringSliceVar(&certConfig.URIs, "uri", []string{}, "uri subject alternative name")
	cmd.Flags().StringVar(&certConfig.ExpireIn, "expire-in", "1 year", "duration of cert validity")
}
