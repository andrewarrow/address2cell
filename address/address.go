package address

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"googlemaps.github.io/maps"
)

func Process(body string) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_GEO")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return
	}

	lines := strings.Split(body, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		fmt.Println(trimmedLine)
		GeocodeAdress(c, trimmedLine)
	}
}

func GeocodeAdress(c *maps.Client, address string) {
	r := &maps.GeocodingRequest{
		Address: address,
	}

	res, err := c.Geocode(context.Background(), r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res[0].Geometry.Location.Lat, res[0].Geometry.Location.Lng)
}
