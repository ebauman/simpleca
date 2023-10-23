package cert

import (
	"errors"
	"github.com/ebauman/simpleca/parse"
	"github.com/ebauman/simpleca/tls"
	"github.com/manifoldco/promptui"
	"reflect"
)

// signUI uses the field names from the certConfig struct to prompt for accompanying user input and signs a new cert
// with the resulting data
func signUI() (err error) {
	certConfig := tls.NewCertConfig()
	certConfig.Path = caPath

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
		"CommonName": promptui.Prompt{
			Label:    "Common Name",
			Default:  "www.simpleca.org",
			Validate: parse.ValidateDNSNames,
		},
	}
	fields := reflect.TypeOf(tls.CertConfig{})

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Field(i).Name
		prompt, ok := prompts[fieldName]
		if !ok {
			continue
		}
		res, err := prompt.Run()
		if err != nil {
			return err
		} else if res == "" {
			continue
		}
		certConfig.SetField(fieldName, res)
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
