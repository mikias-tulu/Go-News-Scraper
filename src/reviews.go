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

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
			var name string
			name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			var stars string
			stars = h.ChildText("span.a-icon-alt")
			var price string
			price = h.ChildText("span.a-price-whole")

			fmt.Println("ProductName: ", name)
			fmt.Println("Ratings: ", stars)
			fmt.Println("Price: ", price)

		})
	})
	c.Visit("https://www.amazon.in/s?k=keyboard")

}
