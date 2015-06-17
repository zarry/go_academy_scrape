package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func scrape(baseURL string, user string) {
	var totalActivity int64
	var dayStreak int64
	doc, err := goquery.NewDocument(baseURL + user)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	doc.Find(".grid-row.profile-time").Each(func(i int, start *goquery.Selection) {
		start.Find("h3").Each(func(i int, it *goquery.Selection) {
			if count == 0 {
				totalActivity, err = strconv.ParseInt(it.Text(), 0, 10)
			} else {
				dayStreak, err = strconv.ParseInt(it.Text(), 0, 10)
			}
			count = count + 1
		})
	})
	print(user, totalActivity, dayStreak)
}

func print(user string, totalActivity int64, dayStreak int64) {
	fmt.Println("User: " + user)
	fmt.Println("---------------------")
	fmt.Println("Day Streak: " + strconv.FormatInt(dayStreak, 10))
	fmt.Println("Total Acitivty: " + strconv.FormatInt(totalActivity, 10))
	fmt.Println("")
}

func main() {
	fmt.Println("")
	fmt.Println("################################")
	fmt.Println("## The Code Academy Challenge ##")
	fmt.Println("################################")
	fmt.Println("")
	baseURL := "http://www.codecademy.com/"
	users := [3]string{"Zarou", "Zarouu", "Zarry"}
	for _, user := range users {
		scrape(baseURL, user)
	}

}
