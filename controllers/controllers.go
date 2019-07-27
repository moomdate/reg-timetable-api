package controllers

import (
	"strings"
	"github.com/gocolly/colly"
	"github.com/labstack/echo"
	"net/http"
)
// CourseDetail ...
type CourseDetail struct {
	Name     string `json:"name"`
	Group    string `json:"group"`
	CourseID string `json:"course_id"`
	Version  string `json:"_version"`
}

// var ...
// url ...
// asd ...
// var A = "initial"
const (
	Baseroot   = "body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody"
	TermDetail = "body > table > tbody > tr:nth-child(1) > td:nth-child(3) > table:nth-child(3) > tbody > tr:nth-child(7) > td:nth-child(2) > font > font"
	// URL        = "http://reg5.sut.ac.th/registrar/learn_time.asp?studentid=15917273&f_cmd=2"
)

// Course ...
type Course struct {
	Acadyear string         `json:"_acadyear"`
	Semester string         `json:"_semester"`
	Data     []CourseDetail `json:"course_list"`
}
// RenderCourseDetail ...
func RenderCourseDetail(ce echo.Context)  error {
	
	stdid := ce.Param("stdid")
	acadyears := ce.Param("acadyear")
	semesters := ce.Param("semester")
	// sli := string(stdid[0])
	// stdCheck := strings.Split(stdid, string([]rune(stdid)[0]))
	
	if strings.ToUpper(stdid[0:1]) == "B" {
		stdid = "1"+stdid[1:]
	} else if strings.ToUpper(stdid[0:1]) == "M" {
		stdid = "2"+stdid[1:]
	} else {
		stdid = "3"+stdid[1:]
	}
	// fmt.Println(stdid)
	URL := "http://reg5.sut.ac.th/registrar/learn_time.asp?f_cmd=2&studentid="+ stdid + "&acadyear=" +acadyears+ "&maxsemester=3&rnd=43673.6771527778&firstday=22/7/2562&semester="+semesters

	// srt := stdid+acadyears+semesters
	// fmt.Println(srt)
	// return c.String(http.StatusOK, srt)
	
	c := colly.NewCollector()
	var cid, gid, cname,ver string
	// var courseD CourseDetail
	// var coursess Course
	var Data []CourseDetail
	c.OnHTML("body > table > tbody > tr:nth-child(1) > td:nth-child(3) > font > b > div > table > tbody > tr > td > table > tbody > tr > td > font > table > tbody > tr > td > table > tbody > tr > td > table > tbody > tr", func(e *colly.HTMLElement) {
		// fmt.Println(e.Index)

		if e.Index != 0 {
			e.ForEach("td", func(_ int, el *colly.HTMLElement) {
				// fmt.Println(el.Index, e.Index)
				if el.Index == 0 {
					// fmt.Println(el.Text)
					// cid = el.Text
					cid = strings.TrimSpace(el.Text)
					datacid := strings.Split(cid, "-")
					cid = datacid[0]
					ver = datacid[1]
				}
				if el.Index == 1 {
					// fmt.Println(el.Text)
					for i := 0; i < len(el.Text); i++ {
						if int(el.Text[i]) > 160 {
							// fmt.Println(string([]rune(el.Text)[i])) // UTF-8
							s := strings.Split(el.Text, string([]rune(el.Text)[i]))
							cname = s[0]
							// fmt.Println(s)
							break
						}
					}
					// cname = el.Text
				}
				if el.Index == 2 {
					// fmt.Println(el.Text)
					gid = el.Text
				}
				// es :=
			})
			courseD := CourseDetail{
				Name:     cname,
				Group:    gid,
				CourseID: cid,
				Version: ver,
			}

			Data = append(Data, courseD)

		}

	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("heard", r.Headers)
		r.ResponseCharacterEncoding = "charset=utf-8"
	})

	c.Visit(URL)
	// fmt.Println(Data)
	couses := &Course{
		Acadyear: acadyears,
		Semester: semesters,
		Data:     Data,
	}
	return ce.JSON(http.StatusOK, couses)
}
