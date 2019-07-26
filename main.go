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
	Name  string   `json:"name"`
	Group int      `json:"group"`
	Time  []string `json:"time"`
}

// Course in register this term
type Course struct {
	CourseList []string `json:"course_list"`
	Data       []CourseDetail
}

func main() {
	c := colly.NewCollector()
	url := "http://reg5.sut.ac.th/registrar/learn_time.asp?studentid=15917273&f_cmd=2&studentname=%B9%D2%C2+%C0%D9%E4%C1%C2%EC+%A8%D1%B9%B7%C3%EC%C0%D9%A7%D2"
	var couse []string
	// Find and visit all links
	c.OnHTML("body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// "body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr:nth-child(2) > td:nth-child(2) > font"
		// with ,_ : =strconv.Atoi()
		// che := e.ChildText("tr:nth-child(3) > td:nth-child(2) > font")
		// fmt.Println(che)
		// str := strings.TrimSpace(e.Text)
		// fmt.Println(len(str))
		// strs := strings.Split(str, " ")
		//
		// fmt.Println(strs)
		// for i := 0; i < len(strs); i++ {
		// 	fmt.Println(strs[i])
		// }

		e.ForEach("table > tbody > tr > td:nth-child(2) > font", func(_ int, el *colly.HTMLElement) {
			if el.Text != "ชื่อรายวิชา" {
				// fmt.Println(el.Text)
				couse = append(couse, el.Text)

			}
		})

	})
	c.OnHTML("body > table > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(3) > tbody > tr:nth-child(7) > td:nth-child(2) > font > font", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		es := strings.Split(e.Text, "/")
		// es2 := strings.Split(es[1], "/")
		fmt.Println(es[0]) // acadyear
		fmt.Println(es[1]) // semester
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("heard", r.Headers)
		r.ResponseCharacterEncoding = "charset=utf-8"

		// r.
	})

	c.Visit(url)
	courseD := CourseDetail{
		Name:  "keng",
		Group: 1,
		Time:  []string{"1-2", "2-3"},
	}
	// var cd []CourseDetail
	// fmt.Println(len(couse))
	// cd = append(cd, courseD)
	course := &Course{
		CourseList: couse,
		Data:       append([]CourseDetail{}, courseD),
	}
	// coursList, _ := json.Marshal(course)
	// fmt.Println(string(coursList))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, course)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
