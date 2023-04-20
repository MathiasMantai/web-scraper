package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"os"
)

func main() {
	var webpage string

	fmt.Println("Pls enter a url: ")
	fmt.Scan(&webpage)

	response, err := http.Get(webpage)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Website could not be scraped")
		//log errors
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	title := document.Find("title").Text()
	content, err := document.Find("html").Html()
	if err != nil {
		panic(err)
	}
	os.WriteFile(title + ".html", []byte(content), 0644)
}