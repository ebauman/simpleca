package parse

import "net/url"

func ParseURIs(uris []string) ([]*url.URL, error) {
	var urls = make([]*url.URL, 0)
	for _, rawURI := range uris {
		u, err := url.Parse(rawURI)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	return urls, nil
}
