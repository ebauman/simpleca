package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vltraheaven/simpleca/cmd/ca"
	"github.com/vltraheaven/simpleca/cmd/cert"
	"github.com/vltraheaven/simpleca/cmd/interactive"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "simpleca",
	Short: "simpleca: hassle-free ca-signed certificate generation",
	Long: `simpleca is used to generate certificate authorities and sign certificates.
It is intended to be hassle-free, relying on defaults as much as possible but allowing for configuration. 
It is also intended to be for developers or for testing purposes. Expressly not recommended for production.`,
}

func init() {
	rootCmd.AddCommand(ca.CAcmd)
	rootCmd.AddCommand(cert.Certcmd)
	rootCmd.AddCommand(interactive.Interactivecmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
