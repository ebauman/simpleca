package cert

import (
	"fmt"
	"github.com/ebauman/simpleca/cmd/interactive"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net"
	"reflect"
	"strings"
)

func init() {
	interactive.Interactivecmd.AddCommand(Certprompt)
}

var Certprompt = &cobra.Command{
	Use:   "cert",
	Short: "Interactive Certificate management",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Choose operation",
			Items: []string{"init"},
		}
		_, res, err := prompt.Run()
		if err != nil {
			fmt.Errorf("an error has occured: %w", err)
		}
		switch res {
		case "init":
			err = certInit()
		}
		if err != nil {
			fmt.Errorf("an error has occured: %w", err)
		}
	},
}

func certInit() error {
	type results struct {
		Passphrase, Country, State, Locality, Organization, OU, Expiry string
		DNS, URI, Email                                                []string
		IP                                                             []net.IP
	}
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
			Default: "",
		},
		"Organization": promptui.Prompt{
			Label:   "Organization",
			Default: "SimpleCA Ltd.",
		},
		"OU": promptui.Prompt{
			Label:   "Organizational unit",
			Default: "SimpleCA Security",
		},
		"IP": promptui.Prompt{
			Label: "Comma-delimited list of IP Addresses",
		},
		"DNS": promptui.Prompt{
			Label: "Comma-delimited list of DNS Names",
		},
		"Email": promptui.Prompt{
			Label: "Comma-delimited list of Email Addresses",
		},
		"URI": promptui.Prompt{
			Label: "Comma-delimited list of URI Subject Names",
		},
		"Expiry": promptui.Prompt{
			Label:   "Expire-in",
			Default: "1 year",
		},
	}
	promptResults := results{}
	fields := reflect.TypeOf(promptResults)

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
		case "ip":
			var netIP []net.IP
			for _, ip := range strings.Split(res, ",") {
				netIP = append(netIP, net.ParseIP(ip))
			}
			reflect.ValueOf(&promptResults).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(netIP))
		case "dns", "email", "uri":
			reflect.ValueOf(&promptResults).Elem().FieldByName(fieldName).Set(
				reflect.ValueOf(slicer(res)))
		default:
			reflect.ValueOf(&promptResults).Elem().FieldByName(fieldName).SetString(res)
		}
	}
	fmt.Println(promptResults)
	return nil
}
