package ca

import (
	"errors"
	"fmt"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/parse"
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"reflect"
	"strings"
)

// initUI uses the field names from the certConfig struct to prompt for accompanying user input, then uses the input to
// initialize a new CA
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
			Default: "Electric",
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
			Label:    "Comma-delimited list of IP Addresses",
			Validate: parse.ValidateIPAddresses,
			Default:  "0.0.0.0",
		},
		"DNSNames": promptui.Prompt{
			Label:    "Comma-delimited list of DNS Names",
			Validate: parse.ValidateDNSNames,
			Default:  "ca.simpleca.org",
		},
		"EmailAddresses": promptui.Prompt{
			Label:    "Comma-delimited list of Email Addresses",
			Validate: parse.ValidateEmailAddresses,
			Default:  "noreply@example.org",
		},
		"URIs": promptui.Prompt{
			Label:    "Comma-delimited list of URI Subject Names",
			Validate: parse.ValidateURIs,
			Default:  "https://ca.simpleca.org",
		},
		"ExpireIn": promptui.Prompt{
			Label:    "Duration of CA validity",
			Default:  "1 year",
			Validate: parse.ValidateDuration,
		},
		"Name": promptui.Prompt{
			Label:   "CA Name",
			Default: "default",
		},
		"Path": promptui.Prompt{
			Label:    "Path to CA store directory",
			Default:  caPath,
			Validate: parse.ValidatePath,
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

		switch strings.ToLower(fieldName) {
		case "ipaddresses":
			ips := parse.ConvertToIPSlice(parse.ConvertToStringSlice(res))
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(ips))
		case "dnsnames", "emailaddresses", "uris":
			v := parse.ConvertToStringSlice(strings.ToLower(res))
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(v))
		default:
			reflect.ValueOf(certConfig).Elem().FieldByName(fieldName).SetString(strings.TrimSpace(res))
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
