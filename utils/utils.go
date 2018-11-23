package utils

import "net/url"

func MakeUrl(shortUrl string) string {
	u, _ := url.Parse("http://localhost:8080")
	u.Path = shortUrl
	return u.String()
}
