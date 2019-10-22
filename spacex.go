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

	message := fmt.Sprintf("Most Recent Launch\nFlight Number: %f\nMission Name: %s\nLaunch Year: %s\nLaunch Success %v", launch.FlightNumber, launch.MissionName, launch.LaunchYear, launch.LaunchSuccess)

	return message
}
