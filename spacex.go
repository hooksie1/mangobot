package p

import (
	"fmt"
	"log"

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
	c, err := spacex.NewWithRateLimit(50)
	if err != nil {
		log.Println(err)
	}

	object, err := c.GetObjectInfo("roadster")
	if err != nil {
		log.Println(err)
	}

	message := fmt.Sprintf("Name: %s\nSpeed: %f MPH\nDistance from Earth: %f Miles", object.Name, object.SpeedMph, object.EarthDistanceMi)

	return message
}
