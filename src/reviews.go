package main

import (
	//"encoding/csv"
	"fmt"
	//"log"
	//"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	c.Visit("https://www.amazon.in/s?k=keyboard")
}
