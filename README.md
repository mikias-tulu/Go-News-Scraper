# GoProductReviewScraper

GoProductReviewScraper is a web scraping project in GoLang that fetches and stores product reviews from popular e-commerce websites. Users can specify a product category or keyword to retrieve the latest reviews related to their interests. The extracted data is stored in a structured format for further analysis or presentation.

**Web Scraping with Colly**

This project uses Colly, a powerful web scraping framework in GoLang. Colly provides a simple API for performing web scraping tasks and is well-suited for writing crawlers, scrapers, and spiders.

## Features

- Web Scraping: Extract product reviews using GoLang's HTTP client and HTML parsing libraries.
- User Input: Select a product category or enter custom keywords to fetch relevant reviews.
- Error Handling: Gracefully handle connectivity issues and invalid user input.
- Data Storage: Store extracted review data in a structured format (JSON/CSV).
- Caching: Reduce requests to the website by implementing a caching mechanism.
- Pagination Handling: Fetch reviews from multiple pages with pagination support.
- Asynchronous Scraping: Improve performance using asynchronous web scraping.
- Respectful Scraping: Comply with the website's terms of service and scraping policies.

## Usage

1. Clone the repository: ``` git clone https://github.com/mikias-tulu/Go-ProductReview-Scraper.git ```
 2. Navigate to the project directory: ``` cd Go-ProductReview-Scraper ```
 3. Install dependencies.
 4. Run the application: ``` go run main.go```
 5. Follow the on-screen prompts to specify the category or topic for news articles.

## Contributing

Contributions are welcome! If you have suggestions, bug reports, or want to add new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
