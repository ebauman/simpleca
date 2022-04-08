package parse

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"path/filepath"
	"strings"
)

// Functions for validating user input

func ValidateIPAddresses(s string) (err error) {
	if s == "" {
		return
	}
	var invalid []string
	for _, t := range strings.Split(s, ",") {
		t = strings.TrimSpace(t)
		if ok := govalidator.IsIP(t); !ok {
			invalid = append(invalid, t)
		}
	}
	if invalid != nil {
		err = fmt.Errorf("invalid IP addresses: %s", strings.Join(invalid, ", "))
	}
	return
}

func ValidateDNSNames(s string) (err error) {
	var invalid []string
	for _, t := range strings.Split(s, ",") {
		t = strings.TrimSpace(t)
		if strings.HasPrefix(t, "*") {
			t = strings.TrimLeft(t, "*.")
		}

		if ok := govalidator.IsDNSName(t); !ok {
			invalid = append(invalid, t)
		}
	}
	if invalid != nil {
		err = fmt.Errorf("invalid DNS names: %s", strings.Join(invalid, ", "))
	}
	return
}

func ValidateEmailAddresses(s string) (err error) {
	var invalid []string
	for _, t := range strings.Split(s, ",") {
		t = strings.ToLower(strings.TrimSpace(t))
		if ok := govalidator.IsEmail(t); !ok {
			invalid = append(invalid, t)
		}
	}
	if invalid != nil {
		err = fmt.Errorf("invalid email addresses: %s", strings.Join(invalid, ", "))
	}
	return
}

func ValidateURIs(s string) (err error) {
	var invalid []string
	for _, t := range strings.Split(s, ",") {
		t = strings.ToLower(strings.TrimSpace(t))
		if ok := govalidator.IsURL(t); !ok {
			invalid = append(invalid, t)
		}
	}
	if invalid != nil {
		err = fmt.Errorf("invalid URIs: %s", strings.Join(invalid, ", "))
	}
	return
}

func ValidatePath(s string) (err error) {
	_, err = filepath.Abs(strings.TrimSpace(s))
	return
}

func ValidateDuration(s string) (err error) {
	_, err = ConvertDuration(s)
	return
}
