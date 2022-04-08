package ca

import (
	"github.com/spf13/cobra"
	"github.com/vltraheaven/simpleca/tls"
	"os"
)

var capath string
var homeDir, _ = os.UserHomeDir()

func init() {
	CAcmd.PersistentFlags().StringVar(&capath, "path", homeDir+"/.simpleca",
		"path where certificate authorities are stored")
}

var certConfig = &tls.CertConfig{}

var CAcmd = &cobra.Command{
	Use:   "ca",
	Short: "operations on certificate authorities",
}
