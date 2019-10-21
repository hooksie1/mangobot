// Package p contains an HTTP Cloud Function.
package p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

var token = os.Getenv("TOKEN")
var port = os.Getenv("PORT")
var url = os.Getenv("URL")

func httpRequest(u string) []byte {
	req, err := http.Get(u)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	return body
}

func getStatus() string {
	log.Println("getting status")
	format := "2006-01-02T15:04:05.000Z"
	date := time.Now()

	var user mangoUser

	req, err := http.Get("https://mangolassi.it/api/user/dustinb3403")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &user)

	printBan := "Not banned"

	if user.Banned {
		then, err := time.Parse(format, user.BannedUntilReadable.Format(format))
		if err != nil {
			log.Println(err)
		}

		diff := then.Sub(date)

		hours := diff.Hours()

		remain := math.Mod(hours, 24)

		days := int(diff.Hours() / 24)

		printBan = fmt.Sprintf("Still banned. Unbanned in %v days and %v hours.", days, remain)
	}

	return printBan

}

func getMetrics() string {
	var categories Categories
	var users Users
	var posts int
	var topics int

	cats := httpRequest("https://mangolassi.it/api/categories")

	totUsers := httpRequest("https://mangolassi.it/api/users")

	err := json.Unmarshal(cats, &categories)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(totUsers, &users)
	if err != nil {
		log.Println(err)
	}

	for _, v := range categories.Categories {
		posts += v.TotalPostCount
		topics += v.TotalTopicCount
	}

	message := fmt.Sprintf("Users: %v\nPosts: %v\nTopics: %v", users.UserCount, float64(posts), float64(topics))

	return message
}

func sendMessage(m Message, s string) {
	log.Println("setting up message")
	var response Response

	response.ChatID = m.Message.Chat.ID

	botURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	response.Text = s

	var body []byte

	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))

	req, err := http.Post(botURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

}

func Bot(w http.ResponseWriter, r *http.Request) {

	log.Println("started")

	var message Message

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &message)

	log.Println(string(body))
	log.Println(message)

	if message.Message.Text == "/dustin" || message.Message.Text == "/dustin@mangobannedbot" {
		status := getStatus()
		sendMessage(message, status)
	}

	if message.Message.Text == "/mangometrics" || message.Message.Text == "/mangometrics@mangobannedbot" {
		metrics := getMetrics()
		sendMessage(message, metrics)
	}

	if message.Message.Text == "/recent" || message.Message.Text == "/recent@mangobannedbot" {
		recent := getRecent()
		sendMessage(message, recent)
	}

}
