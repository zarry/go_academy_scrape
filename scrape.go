package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// User struct holds dayStreak and totalActivity from scrape
type User struct {
	dayStreak, totalActivity int64
}

var usersString string
var baseURL = "http://www.codecademy.com/"
var results = make(map[string]User)

func init() {
	flag.StringVar(&usersString, "users", "", "A comma seperated list of users to scrape")
}

func scrape(baseURL string, user string) (dayStreak, totalActivity int64) {
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
	return dayStreak, totalActivity
}

func print(user string, totalActivity int64, dayStreak int64) {
	fmt.Println("User: " + user)
	fmt.Println("---------------------")
	fmt.Println("Day Streak: " + strconv.FormatInt(dayStreak, 10))
	fmt.Println("Total Acitivty: " + strconv.FormatInt(totalActivity, 10))
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

	users := strings.Split(usersString, ",")

	for _, user := range users {
		trimmedUser := strings.TrimSpace(user)
		streak, activity := scrape(baseURL, trimmedUser)
		results[trimmedUser] = User{streak, activity}
	}

	// Results
	for user, data := range results {
		print(user, data.totalActivity, data.dayStreak)
	}

}
