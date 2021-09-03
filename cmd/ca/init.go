package ca

import (
	"fmt"
	"github.com/ebauman/simpleca/cli"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	CAcmd.AddCommand(initCmd)

	cli.LoadCertFlags(certConfig, initCmd)

	initCmd.Flags().StringVar(&certConfig.Name, "name", "default", "vanity name for the certificate")
	initCmd.Flags().StringVar(&certConfig.CommonName, "common-name", "SimpleCA Root Certificate Authority",
		"common name")
	initCmd.Flags().StringVar(&certConfig.Path, "path", "$HOME/.simpleca",
		"path where certificate will be put. subdirectory will be created to match vanity name")
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "create a new certificate authority",
	RunE: func(cmd *cobra.Command, args []string) error {
		if certConfig.Path == "$HOME/.simpleca" {
			certConfig.Path = os.Getenv("HOME") + "/.simpleca"
		}
		if err := file.CheckPath(fmt.Sprintf("%s/%s", certConfig.Path, certConfig.Name)); err != nil {
			return err
		}

		return tls.GenerateCA(certConfig)
	},
}
