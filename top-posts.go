package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//<a href="http://google.com">link</a>.'

func getRecent() Recent {

	var recent Recent

	req, err := http.Get("https://mangolassi.it/api/recent")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &recent)

	return recent

}

func createRecent() string {
	topics := getRecent()

	message := ""
	link := ""

	for i := 0; i < 10; i++ {
		link = fmt.Sprintf("<a href=\"https://mangolassi.it/topic/%v\">%s</a>", topics.Topics[i].Slug, topics.Topics[i].Title)
		message += string(i) + " - " + link + "\n"
	}

	return message
}
