package ca

import (
	"errors"
	"fmt"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"net"
	"reflect"
	"strings"
)

func initUI() (err error) {
	var prompts = map[string]promptui.Prompt{
		"Passphrase": promptui.Prompt{
			Label:   "Passphrase (default: changeme)",
			Mask:    '*',
			Default: "changeme",
		},
		"Country": promptui.Prompt{
			Label:   "Country code",
			Default: "AA",
		},
		"State": promptui.Prompt{
			Label:   "State",
			Default: "Relaxation",
		},
		"Locality": promptui.Prompt{
			Label:   "Locality",
			Default: "",
		},
		"Organization": promptui.Prompt{
			Label:   "Organization",
			Default: "SimpleCA Ltd.",
		},
		"OrganizationalUnit": promptui.Prompt{
			Label:   "Organizational unit",
			Default: "SimpleCA Security",
		},
		"IPAddresses": promptui.Prompt{
			Label: "Comma-delimited list of IP Addresses",
		},
		"DNSNames": promptui.Prompt{
			Label: "Comma-delimited list of DNS Names",
		},
		"EmailAddresses": promptui.Prompt{
			Label: "Comma-delimited list of Email Addresses",
		},
		"URIs": promptui.Prompt{
			Label: "Comma-delimited list of URI Subject Names",
		},
		"ExpireIn": promptui.Prompt{
			Label:   "Duration of CA validity",
			Default: "1 year",
		},
		"Name": promptui.Prompt{
			Label:   "CA Name",
			Default: "default",
		},
		"Path": promptui.Prompt{
			Label:   "Path to CA store directory",
			Default: caPath,
		},
		"CommonName": promptui.Prompt{
			Label:   "Common Name",
			Default: "SimpleCA Root Certificate Authority",
		},
	}

	fields := reflect.TypeOf(tls.CertConfig{})

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Field(i).Name
		prompt := prompts[fieldName]
		res, err := prompt.Run()
		if err != nil {
			return err
		}

		slicer := func(s string) (slice []string) {
			for _, sl := range strings.Split(s, ",") {
				slice = append(slice, sl)
			}
			return
		}

		switch strings.ToLower(fieldName) {
		case "ipaddresses":
			var netIP []net.IP
			for _, ip := range strings.Split(res, ",") {
				netIP = append(netIP, net.ParseIP(ip))
			}
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(netIP))
		case "dnsnames", "emailaddresses", "uris":
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(slicer(res)))
		default:
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).SetString(res)
		}
	}

	confirm := promptui.Prompt{
		Label:     "Confirm",
		IsConfirm: true,
	}

	if _, err = confirm.Run(); err != nil {
		err = errors.New("process aborted")
		return
	}
	if err = file.CheckPath(fmt.Sprintf("%s/%s", certConfig.Path, certConfig.Name)); err != nil {
		return
	}
	return tls.GenerateCA(certConfig)
}
