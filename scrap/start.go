package scrap

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func Run() {

	c := colly.NewCollector()

	c.OnHTML(".table tbody tr", func(e *colly.HTMLElement) {
		rank := e.ChildText(".rankings-block_banner--pos, .table-body_cell--position")
		team := e.ChildText(".u-hide-phablet a")
		matches := e.ChildText(".table-body__cell.u-center-text")
		points := e.ChildText(".table-body__cell.u-center-text:nth-child(3)")
		rating := e.ChildText(".table-body__cell.u-text-right.rating")
		// fmt.Printf("%s\t%-40s\t%s\t%s\t%s\n", rank, team, matches, points, rating)
		fmt.Println(rank, "rank")
		fmt.Println(team, "team")
		fmt.Println(points, "points")
		fmt.Println(matches, "matches")
		fmt.Println(rating, "rating")

	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("\n Response received: ", r.StatusCode, " ends \n ")
	})

	c.Visit("https://www.icc-cricket.com/rankings/mens/overview")

	// Wait for the collector to finish
	c.Wait()
}
