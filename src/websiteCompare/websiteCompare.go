package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

//https://www.golangprograms.com/how-to-extract-text-from-between-html-tag-using-regular-expressions-in-golang.html

func main() {
	url := os.Args[1]
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
	re := regexp.MustCompile(`<p.*?>(.*)</p>`)
	re2 := regexp.MustCompile(`<h1.*?>(.*)</h1>`)
	re3 := regexp.MustCompile(`<h2.*?>(.*)</h2>`)
	re4 := regexp.MustCompile(`<h3.*?>(.*)</h3>`)
	re5 := regexp.MustCompile(`<h4.*?>(.*)</h4>`)
	re6 := regexp.MustCompile(`<h5.*?>(.*)</h5>`)
	re7 := regexp.MustCompile(`<h6.*?>(.*)</h6>`)
	re8 := regexp.MustCompile(`<a.*?>(.*)</a>`)
	bodyText := re.FindAllStringSubmatch(html, -1)+re2.FindAllStringSubmatch(html, -1)+re3.FindAllStringSubmatch(html, -1)+re4.FindAllStringSubmatch(html, -1)+re5.FindAllStringSubmatch(html, -1)+re6.FindAllStringSubmatch(html, -1)+re7.FindAllStringSubmatch(html, -1)+re8.FindAllStringSubmatch(html, -1)
	fmt.Println(bodyText)
}
