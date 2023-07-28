package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	csvFile, err := os.Create("scraped_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Product Name", "Ratings", "Price"})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
			var name string = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			var stars string = h.ChildText("span.a-icon-alt")
			var price string = h.ChildText("span.a-price-whole")

			fmt.Println("ProductName: ", name)
			fmt.Println("Ratings: ", stars)
			fmt.Println("Price: ", price)

			// Write data to CSV
			writer.Write([]string{name, stars, price})
		})
	})

	// Scrape data from multiple pages of Amazon search results for keyboards
	numPages := 10 // set value to scrape more or fewer pages

	baseURL := "https://www.amazon.in/s?k=keyboard&page="
	for page := 1; page <= numPages; page++ {
		url := baseURL + fmt.Sprintf("%d", page)
		c.Visit(url)
	}
}
