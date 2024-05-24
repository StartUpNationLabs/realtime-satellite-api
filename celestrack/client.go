package celestrack

import (
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/spacetrack"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

type Client struct {
	client *resty.Client
	conf   Conf
}
type Conf struct {
	baseUrl string
}

func New() Client {
	c := Client{}
	c.client = resty.New()
	if url := os.Getenv("CELESTRACK_URL"); url != "" {
		log.Info("Using CELESTRACK_URL: ", url)
		c.conf.baseUrl = url
	} else {
		c.conf.baseUrl = "https://celestrak.com/NORAD/elements"

	}

	return c
}

func extractNoradIdFromTle(line2 string) string {
	// remove the 0 in the beginning
	dirtyNorad := strings.Split(line2, " ")[1]
	return strings.TrimLeft(dirtyNorad, "0")

}

func (c *Client) Scrap() (map[string]spacetrack.TLE, error) {
	// Get the HTML
	resp, err := c.client.R().Get(c.conf.baseUrl + "/")
	if err != nil {
		return nil, err
	}

	// Find the TLE files
	re := regexp.MustCompile(`gp\.php\?GROUP=(.+?)&`)
	matches := re.FindAllStringSubmatch(resp.String(), -1)
	data := make(map[string]spacetrack.TLE)
	for _, match := range matches {
		group := match[1]
		log.Info("Fetching Group: ", group)

		// Get the TLE file
		resp, err := c.client.R().Get(c.conf.baseUrl + "/gp.php?GROUP=" + group + "&FORMAT=tle")
		if err != nil {
			return nil, err
		}
		lines := strings.Split(resp.String(), "\n")
		for i := 0; i < len(lines); i += 3 {

			norad := extractNoradIdFromTle(lines[i+2])
			tle, ok := data[norad]
			if !ok {
				tle = spacetrack.TLE{
					TLE_LINE0:    lines[i],
					TLE_LINE1:    lines[i+1],
					TLE_LINE2:    lines[i+2],
					NORAD_CAT_ID: norad,
					OBJECT_NAME:  strings.TrimSpace(lines[0]),
					Group:        []string{group},
				}
			} else {
				tle.Group = append(tle.Group, group)
			}
			data[norad] = tle
		}
	}
	return data, nil
}
