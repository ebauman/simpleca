package ca

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vltraheaven/simpleca/cli"
	"github.com/vltraheaven/simpleca/file"
	"github.com/vltraheaven/simpleca/tls"
)

func init() {
	CAcmd.AddCommand(initCmd)

	cli.LoadCertFlags(certConfig, initCmd)

	initCmd.Flags().StringVar(&certConfig.Name, "name", "default", "vanity name for the certificate")
	initCmd.Flags().StringVar(&certConfig.CommonName, "common-name", "SimpleCA Root Certificate Authority",
		"common name")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a new certificate authority",
	RunE: func(cmd *cobra.Command, args []string) error {
		certConfig.Path = capath
		if err := file.CheckPath(fmt.Sprintf("%s/%s", certConfig.Path, certConfig.Name)); err != nil {
			return err
		}

		return tls.GenerateCA(certConfig)
	},
}
