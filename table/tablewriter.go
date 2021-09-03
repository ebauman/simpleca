package table

import (
	"crypto/x509"
	"fmt"
	"github.com/liggitt/tabwriter"
	"io"
	"strings"
)

func WriteCAList(out io.Writer, cas map[string]*x509.Certificate) error {
	tw := tabwriter.NewWriter(out, 6, 4, 3, ' ', tabwriter.RememberWidths)
	defer tw.Flush()

	headers := []string{"NAME", "VALID FROM", "VALID UNTIL", "EXPIRES IN"}
	_, err := fmt.Fprintf(tw, "%s\n", strings.Join(headers, "\t"))
	if err != nil {
		return err
	}

	for k, c := range cas {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t\n", k, c.NotBefore, c.NotAfter, c.NotAfter.Sub(c.NotBefore))
	}

	return nil
}
