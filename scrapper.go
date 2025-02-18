package main

import (

	//"os"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	Name         string `db:"name"`
	Price        string `db:"price"`
	Review       string `db:"review"`
	DateScrapped string `db:"datescrapped"`
}

func extractSubstring(input string) string {
	brIndex := strings.Index(input, "<br>")
	if brIndex == -1 {
		return "" // Return an empty string if "<br>" is not found
	}
	return input[:brIndex]
}

func webScrap() {
	var items []item
	c := colly.NewCollector()

	c.OnHTML("div[id=search_result_container]", func(h *colly.HTMLElement) {
		item := item{}
		h.ForEach("div.responsive_search_name_combined", func(_ int, h *colly.HTMLElement) {
			h.ForEach("div.col.search_reviewscore.responsive_secondrow", func(_ int, h *colly.HTMLElement) {
				h.ForEach("span.search_review_summary.positive", func(_ int, h *colly.HTMLElement) {
					linkStr := h.Attr("data-tooltip-html")
					// fmt.Println("Reviews : ")
					out := extractSubstring(linkStr)
					// fmt.Println(out)
					if len(out) > 0 {
						review := out
						item.Review = review
					}
				})
				h.ForEach("span.search_review_summary.mixed", func(_ int, h *colly.HTMLElement) {
					linkStr := h.Attr("data-tooltip-html")
					out := extractSubstring(linkStr)
					if len(out) > 0 {
						review := out
						item.Review = review
					}
				})
				h.ForEach("span.search_review_summary.negative", func(_ int, h *colly.HTMLElement) {
					linkStr := h.Attr("data-tooltip-html")
					out := extractSubstring(linkStr)
					if len(out) > 0 {
						review := out
						item.Review = review
					}
				})
			})
			h.ForEach("span.title", func(_ int, h *colly.HTMLElement) {
				// fmt.Println(h.Text)
				name := h.Text
				item.Name = name
			})
			h.ForEach("div.col.search_price_discount_combined.responsive_secondrow", func(_ int, h *colly.HTMLElement) {
				h.ForEach("div.discount_prices", func(_ int, h *colly.HTMLElement) {
					h.ForEach("div.discount_final_price", func(_ int, h *colly.HTMLElement) {
						price := h.Text
						item.Price = price
					})
				})
				item.DateScrapped = time.Now().Format("2006-01-02")
			})
			items = append(items, item)
		})
		fmt.Println(items)
	})

	c.Visit("https://store.steampowered.com/search/?filter=topsellers")
}
func main() {

	webScrap()

}
