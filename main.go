package main

import (
	"fmt"
	"strings"

	"github.com/adrian-petrov/go-htmlparser/htmlparser"
)

func main() {
	htmlFile := htmlparser.ReadHTMLFile("html/ex2.html")
	htmlReader := strings.NewReader(string(htmlFile))

	links, err := htmlparser.ParseLinks(htmlReader)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("Link: %#v", link)
	}
}
