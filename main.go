package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	// err := dotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	e := echo.New()
	e.Static("/", "static")
	e.GET("/api/url", func(c echo.Context) error {
		url := c.QueryParam("url")
		html, err := getPremium(url)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Oops!")
		}
		return c.HTML(http.StatusOK, html)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func getPremium(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", "mySPHUserType=sub; SPHiPlanetDirectoryPro="+os.Getenv("SPHiPlanetDirectoryPro"))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
