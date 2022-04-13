package cert

import (
	"github.com/ebauman/simpleca/tls"
	"github.com/spf13/cobra"
	"os"
)

var certConfig = &tls.CertConfig{}
var homeDir, _ = os.UserHomeDir()
var caName string

func init() {
	Certcmd.PersistentFlags().StringVar(&caName, "ca", "default",
		"name of the ca to use")
	Certcmd.PersistentFlags().StringVar(&certConfig.Path, "ca-path", homeDir+"/.simpleca",
		"path where certificate authorities are stored")
}

var Certcmd = &cobra.Command{
	Use:   "cert",
	Short: "operations on certificates",
}
