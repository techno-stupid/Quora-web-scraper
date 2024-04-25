package main

import (
	"fmt"
	"log"

	"github.com/techno-stupid/Quora-web-scraper/scraper"
)

func main() {
	// URL of the Quora page to scrape
	url := "https://www.quora.com/"

	// Call function to scrape questions and answers
	questions, err := scraper.ScrapeQuora(url)
	if err != nil {
		log.Fatalf("Error scraping Quora: %v", err)
	}

	// Print scraped questions and answers
	for _, q := range questions {
		fmt.Printf("Question: %s\n", q.Title)
		fmt.Println("Answers:")
		for _, ans := range q.Answers {
			fmt.Printf("- %s\n", ans)
		}
		fmt.Println()
	}
}
