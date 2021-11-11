package ca

import (
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var capath string

func init() {
	CAcmd.PersistentFlags().StringVar(&capath, "path", os.Getenv("HOME")+"/.simpleca",
		"path where certificate authorities are stored")
}

var certConfig = &tls.CertConfig{}

var CAcmd = &cobra.Command{
	Use:   "ca",
	Short: "operations on certificate authorities",
}
