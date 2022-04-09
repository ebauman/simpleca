package ca

import (
	"fmt"
	"github.com/vltraheaven/simpleca/file"
	"github.com/vltraheaven/simpleca/tls"
)

func ListCAs(path string) (cas []string, err error) {
	var dirs []string
	dirs, err = file.ListDirectories(caPath)
	if err != nil {
		return
	}

	for _, d := range dirs {
		var certPath = fmt.Sprintf("%s/%s/%s", path, d, tls.CACertFileName)
		var keyPath = fmt.Sprintf("%s/%s/%s", path, d, tls.CAKeyFileName)
		_, _, err = tls.LoadCA(certPath, keyPath)
		if err != nil {
			return
		}
		cas = append(cas, d)
	}
	return
}
