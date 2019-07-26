package main

import (
	"echo"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

// CourseDetail ...
type CourseDetail struct {
	Name     string `json:"name"`
	Group    string `json:"group"`
	CourseID string `json:"course_id"`
}

// var ...
// url ...
// asd ...
// var A = "initial"
const (
	Baseroot   = "body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody"
	TermDetail = "body > table > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(3) > tbody > tr:nth-child(7) > td:nth-child(2) > font > font"
	URL        = "http://reg5.sut.ac.th/registrar/learn_time.asp?studentid=15917273&f_cmd=2&studentname=%B9%D2%C2+%C0%D9%E4%C1%C2%EC+%A8%D1%B9%B7%C3%EC%C0%D9%A7%D2"
)

// body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table:nth-child(2) > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr:nth-child(1) > td:nth-child(3) > font
// body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table:nth-child(2) > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr:nth-child(2) > td:nth-child(3) > font
// body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table:nth-child(2) > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr:nth-child(4) > td:nth-child(3) > font
type Course struct {
	CourseList []string       `json:"course_list"`
	Data       []CourseDetail `json:"data"`
}

func main() {
	// fmt.Println(A)
	c := colly.NewCollector()
	var couse []string
	// Find and visit all links
	c.OnHTML(Baseroot, func(e *colly.HTMLElement) {

		e.ForEach("tr > td:nth-child(2) > font", func(_ int, el *colly.HTMLElement) {
			if el.Text != "ชื่อรายวิชา" {
				// fmt.Println(el.Text)
				couse = append(couse, el.Text)

			}
		})
		// e.ForEach("tr > td > font", func(_ int, el *colly.HTMLElement) {
		// 	fmt.Println(el.Index, el.Text)
		// })
		e.ForEach("tr > td:nth-child(1) > font", func(_ int, el *colly.HTMLElement) {
			if el.Text != "รหัสวิชา" {
				// fmt.Println(el.Text)
				// couse = append(couse, el.Text)

			}
		})

	})
	var cid, gid, cname string
	var courseD CourseDetail
	var coursess Course
	c.OnHTML("body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr", func(e *colly.HTMLElement) {
		// fmt.Println(e.Index)

		if e.Index != 0 {
			e.ForEach("td", func(_ int, el *colly.HTMLElement) {
				// fmt.Println(el.Index, e.Index)
				if el.Index == 0 {
					fmt.Println(el.Text)
					cid = el.Text
				}
				if el.Index == 1 {
					fmt.Println(el.Text)
					cname = el.Text
				}
				if el.Index == 2 {
					fmt.Println(el.Text)
					gid = el.Text
				}
				// es :=
			})
			courseD = CourseDetail{
				Name:     cname,
				Group:    gid,
				CourseID: cid,
			}
			// coursess.append([]CourseDetail{}, courseD)
			coursess = Course{
				Data: append([]CourseDetail{}, courseD),
			}
		}

	})
	c.OnHTML(TermDetail, func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		es := strings.Split(e.Text, "/")
		// es2 := strings.Split(es[1], "/")
		fmt.Println(es[0]) // acadyear
		es2 := strings.Split(es[1], "")
		fmt.Println(es2[1]) // semester
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("heard", r.Headers)
		r.ResponseCharacterEncoding = "charset=utf-8"

		// r.
	})

	c.Visit(URL)

	// coursList, _ := json.Marshal(course)
	// fmt.Println(string(coursList))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, coursess)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
