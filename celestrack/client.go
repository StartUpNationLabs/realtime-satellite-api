package celestrack

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
	"regexp"
	"strings"
)

func extractNoradIdFromTle(line2 string) string {
	// remove the 0 in the beginning
	dirtyNorad := strings.Split(line2, " ")[1]
	return strings.TrimLeft(dirtyNorad, "0")

}

func Scrap() (map[string]spacetrack.TLE, error) {
	url := "https://celestrak.org/NORAD/elements/"
	// Get the HTML
	client := resty.New()
	resp, err := client.R().Get(url)
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
		resp, err := client.R().Get(url + "gp.php?GROUP=" + group + "&FORMAT=tle")
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
