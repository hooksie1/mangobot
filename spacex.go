package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/orcaman/spacex"
)

func getLaunch() string {
	c, err := spacex.NewWithRateLimit(50)
	if err != nil {
		log.Println(err)
	}
	launch, err := c.GetLatestLaunch()
	if err != nil {
		log.Println(err)
	}

	message := fmt.Sprintf("Most Recent SpaceX Launch\nFlight Number: %f\nMission Name: %s\nLaunch Year: %s\nLaunch Success: %v", launch.FlightNumber, launch.MissionName, launch.LaunchYear, launch.LaunchSuccess)

	return message
}

func getRoadster() string {
	var roadster Roadster

	req, err := http.Get("https://api.spacexdata.com/v3/roadster")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &roadster)

	pic := ""
	pics := ""

	stats := fmt.Sprintf("Tesla Roadster\nSpeed: %f\nDistance from Earth: %f, Distance from Mars: %f\n Images:\n", roadster.SpeedMph, roadster.EarthDistanceMi, roadster.MarsDistanceMi)

	for i, v := range roadster.FlickrImages {
		j := i + 1
		pic = fmt.Sprintf("<a href=\"%s\">%d</a>", v, j)
		pics += pic + "\n"
	}

	message := fmt.Sprintf("%s\n%s", stats, pics)

	return message
}
