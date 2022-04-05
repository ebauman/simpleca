package cert

import (
	"errors"
	"github.com/ebauman/simpleca/parse"
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"reflect"
	"strings"
)

func signUI() (err error) {
	var prompts = map[string]promptui.Prompt{
		"Passphrase": promptui.Prompt{
			Label:       "Passphrase (default: changeme)",
			Mask:        '*',
			HideEntered: true,
			Default:     "changeme",
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
			Default:  "simpleca.org,*.simpleca.org",
		},
		"EmailAddresses": promptui.Prompt{
			Label:    "Comma-delimited list of Email Addresses",
			Validate: parse.ValidateEmailAddresses,
			Default:  "noreply@example.org",
		},
		"URIs": promptui.Prompt{
			Label:    "Comma-delimited list of URI Subject Names",
			Validate: parse.ValidateURIs,
			Default:  "https://simpleca.org,http://simpleca.org",
		},
		"ExpireIn": promptui.Prompt{
			Label:    "Duration of cert validity",
			Default:  "1 year",
			Validate: parse.ValidateDuration,
		},
		"Name": promptui.Prompt{
			Label:   "Certificate Name",
			Default: "default-cert",
		},
		"Path": promptui.Prompt{
			Label:    "Path to CA store directory",
			Default:  caPath,
			Validate: parse.ValidatePath,
		},
		"CommonName": promptui.Prompt{
			Label:   "Common Name",
			Default: "www.simpleca.org",
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
	return tls.SignCert(certConfig, caName)
}
