package main

import (
	"github.com/labstack/echo"
	"github.com/dtmkeng/scraping/controllers"
)
func main() {
	// fmt.Println(A)

	// coursList, _ := json.Marshal(course)
	// fmt.Println(string(coursList))
	e := echo.New()
	e.GET("/api/v1/:stdid/:acadyear/:semester", controllers.RenderCourseDetail)
	e.Logger.Fatal(e.Start(":1323"))
}
