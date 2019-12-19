package gohtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//GetURLTitles return html title from pages url
func GetURLTitles(urls ...string) chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := ioutil.ReadAll(response.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
