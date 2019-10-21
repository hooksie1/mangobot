package p

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//<a href="http://google.com">link</a>.'

func getRecent() []string {

	var recent Recent

	messages := []string{}

	req, err := http.Get("https://mangolassi.it/api/recent")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &recent)

	for _, v := range recent.Topics {
		messages = append(messages, v.Slug)
	}

	return messages

}

func createRecent() string {
	messages := getRecent()

	message := ""
	for i, v := range messages {
		message += string(i) + v + "\n"
	}

	return message
}
