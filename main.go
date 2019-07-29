package main

import (
	"fmt"
	"os"

	"github.com/dtmkeng/scraping/controllers"
	"github.com/labstack/echo"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "1323"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}

func main() {
	// fmt.Println(A)

	// coursList, _ := json.Marshal(course)
	// fmt.Println(string(coursList))
	e := echo.New()
	e.GET("/api/v1/:stdid/:acadyear/:semester", controllers.RenderCourseDetail)
	// e.GET("/api/v2/:stdid/:acadyear/:semester", controllers.RenderCourseDetail2)
	// e.GET("/api/v2/stdid/:stdid/acadyear/:acadyear/semester/:semester", controllers.RenderCourseDetail)
	e.Logger.Fatal(e.Start(getPort()))
}
