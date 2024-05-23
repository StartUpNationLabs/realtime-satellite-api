package spacetrack

import (
	"os"
)

type Conf struct {
	Url         string
	LoginUrl    string
	FetchAllUrl string
	tleFile     string
	username    string
	password    string
}

const SpaceTrackUrl = "https://www.space-track.org"
const SpaceTrackLoginUri = "/auth/login"
const SpaceTrackFetchAllUri = "/basicspacedata/query/class/gp/decay_date/null-val/epoch/>now-30/orderby/norad_cat_id/format/json"
const defaultTleFile = "./data/spacetrack.json"

func confFromEnv() Conf {
	spaceTrackConf := Conf{}
	// set SpaceTrackUrl from the environment
	if url := os.Getenv("SPACETRACK_URL"); url != "" {
		spaceTrackConf.Url = url
	} else {
		spaceTrackConf.Url = SpaceTrackUrl
	}
	// set SpaceTrackLoginUrl from the environment
	if loginUrl := os.Getenv("SPACETRACK_LOGIN_URL"); loginUrl != "" {
		spaceTrackConf.LoginUrl = loginUrl
	} else {
		spaceTrackConf.LoginUrl = spaceTrackConf.Url + SpaceTrackLoginUri
	}
	// set SpaceTrackApi from the environment
	if fetchAllUrl := os.Getenv("SPACETRACK_FETCH_ALL_URL"); fetchAllUrl != "" {
		spaceTrackConf.FetchAllUrl = fetchAllUrl
	} else {
		spaceTrackConf.FetchAllUrl = spaceTrackConf.Url + SpaceTrackFetchAllUri
	}

	// set username from the environment
	if username := os.Getenv("SPACETRACK_USERNAME"); username != "" {
		spaceTrackConf.username = username
	}
	// set password from the environment
	if password := os.Getenv("SPACETRACK_PASSWORD"); password != "" {
		spaceTrackConf.password = password
	}
	// set tleFile from the environment
	if tleFile := os.Getenv("SPACETRACK_TLE_FILE"); tleFile != "" {
		spaceTrackConf.tleFile = tleFile
	} else {
		spaceTrackConf.tleFile = defaultTleFile
	}
	return spaceTrackConf
}
