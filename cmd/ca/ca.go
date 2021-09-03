package ca

import (
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
)

var certConfig = &tls.CertConfig{}

var CAcmd  = &cobra.Command{
	Use: "ca",
	Short: "operations on certificate authorities",
}