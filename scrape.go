package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// User struct holds dayStreak and totalActivity from scrape
type User struct {
	dayStreak, totalActivity string
}

var usersString string
var baseURL = "http://www.codecademy.com/"
var results = make(map[string]User)

func init() {
	flag.StringVar(&usersString, "users", "", "A comma seperated list of users to scrape")
}

func scrape(baseURL string, user string) (dayStreak, totalActivity string) {
	doc, err := goquery.NewDocument(baseURL + user)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	doc.Find(".grid-row.profile-time").Each(func(i int, start *goquery.Selection) {
		start.Find("h3").Each(func(i int, it *goquery.Selection) {
			if count == 0 {
				totalActivity = it.Text()
			} else {
				dayStreak = it.Text()
			}
			count = count + 1
		})
	})
	return dayStreak, totalActivity
}

func print(user string, totalActivity, dayStreak string) {
	fmt.Println("User: " + user)
	fmt.Println("---------------------")
	fmt.Println("Day Streak: " + dayStreak)
	fmt.Println("Total Acitivty: " + totalActivity)
	fmt.Println("")
}

func printHeader() {
	fmt.Println("")
	fmt.Println("################################")
	fmt.Println("## The Code Academy Challenge ##")
	fmt.Println("################################")
	fmt.Println("")
}

func main() {
	flag.Parse()
	printHeader()

	for _, user := range strings.Split(usersString, ",") {
		trimmedUser := strings.TrimSpace(user)
		streak, activity := scrape(baseURL, trimmedUser)
		results[trimmedUser] = User{streak, activity}
	}

	for user, data := range results {
		print(user, data.totalActivity, data.dayStreak)
	}

}
