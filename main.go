package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"os"
	"encoding/json"
)

//main programm
func main() {
	var webpage string

	fmt.Println("Pls enter a url: ")
	fmt.Scan(&webpage)

	var linkYesNo string
	fmt.Println("Do you also want to scrape all of the links separately?")
	fmt.Scan(&linkYesNo)

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
	if(strings.ToUpper(linkYesNo) == "Y") {
		getLinks(document)
	}

	title := document.Find("title").Text()
	content, err := document.Find("html").Html()
	if err != nil {
		panic(err)
	}

	//create html file
	os.WriteFile(title + ".html", []byte(content), 0644)
}

//scrape all links from html document
func getLinks(document *goquery.Document) {
	title := document.Find("title").Text()
	var links []string

	//find all links and add to slice
	document.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		link, exists := item.Attr("href")
		if exists {
			links = append(links, link)
		}
	})

	//creates byte array for json file
	file, _ := json.MarshalIndent(links, "", " ")

	//write to file
	os.WriteFile("Links_" + title + ".json", file, 0644)
}