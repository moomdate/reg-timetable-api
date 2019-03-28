package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	url := "http://reg5.sut.ac.th/registrar/learn_time.asp?studentid=15917273&f_cmd=2&studentname=%B9%D2%C2+%C0%D9%E4%C1%C2%EC+%A8%D1%B9%B7%C3%EC%C0%D9%A7%D2"
	// Find and visit all links
	c.OnHTML("body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr:nth-child(2)", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))

		str := strings.TrimSpace(e.Text)
		strs := strings.Split(str, " ")

		for i := 0; i < len(strs); i++ {
			fmt.Println(strs[i])
		}

	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("heard", r.Headers)
		r.ResponseCharacterEncoding = "charset=utf-8"

		// r.
	})

	c.Visit(url)
}
