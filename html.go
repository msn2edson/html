package html

import (
	"io"
	"net/http"
	"regexp"
)

// Titulo obtem o titulo de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) // Importante: foi criado a funcao e agora estou invocando
	}
	return c
}
