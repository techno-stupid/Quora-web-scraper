package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeQuora scrapes questions and answers from a Quora page
func ScrapeQuora(url string) ([]QuoraQuestion, error) {
	// Make HTTP GET request
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	// Extract questions and answers
	var questions []QuoraQuestion
	doc.Find(".question_link").Each(func(i int, s *goquery.Selection) {
		questionTitle := s.Text()
		var answers []string
		// Find the sibling div containing answers
		siblings := s.Parent().Next().Find(".ui_qtext_rendered_qtext")
		siblings.Each(func(i int, s *goquery.Selection) {
			answer := strings.TrimSpace(s.Text())
			if answer != "" {
				answers = append(answers, answer)
			}
		})
		// Append question and answers to the list
		questions = append(questions, QuoraQuestion{Title: questionTitle, Answers: answers})
	})

	return questions, nil
}
