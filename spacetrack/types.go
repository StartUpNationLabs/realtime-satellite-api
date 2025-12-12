package spacetrack

import "github.com/joshuaferrara/go-satellite"

// Credentials  is a struct that holds the username and password of the user
type Credentials struct {
	Identity string
	Password string
}
type LoggedInCredentials struct {
	spacecraftCookie string
	chocolatechip    string
}

type TLE struct {
	CCSDS_OMM_VERS      string   `json:"CCSDS_OMM_VERS"`
	COMMENT             string   `json:"COMMENT"`
	CREATION_DATE       string   `json:"CREATION_DATE"`
	ORIGINATOR          string   `json:"ORIGINATOR"`
	OBJECT_NAME         string   `json:"OBJECT_NAME"`
	OBJECT_ID           string   `json:"OBJECT_ID"`
	CENTER_NAME         string   `json:"CENTER_NAME"`
	REF_FRAME           string   `json:"REF_FRAME"`
	TIME_SYSTEM         string   `json:"TIME_SYSTEM"`
	MEAN_ELEMENT_THEORY string   `json:"MEAN_ELEMENT_THEORY"`
	EPOCH               string   `json:"EPOCH"`
	MEAN_MOTION         string   `json:"MEAN_MOTION"`
	ECCENTRICITY        string   `json:"ECCENTRICITY"`
	INCLINATION         string   `json:"INCLINATION"`
	RA_OF_ASC_NODE      string   `json:"RA_OF_ASC_NODE"`
	ARG_OF_PERICENTER   string   `json:"ARG_OF_PERICENTER"`
	MEAN_ANOMALY        string   `json:"MEAN_ANOMALY"`
	EPHEMERIS_TYPE      string   `json:"EPHEMERIS_TYPE"`
	CLASSIFICATION_TYPE string   `json:"CLASSIFICATION_TYPE"`
	NORAD_CAT_ID        string   `json:"NORAD_CAT_ID"`
	ELEMENT_SET_NO      string   `json:"ELEMENT_SET_NO"`
	REV_AT_EPOCH        string   `json:"REV_AT_EPOCH"`
	BSTAR               string   `json:"BSTAR"`
	MEAN_MOTION_DOT     string   `json:"MEAN_MOTION_DOT"`
	MEAN_MOTION_DDOT    string   `json:"MEAN_MOTION_DDOT"`
	SEMIMAJOR_AXIS      string   `json:"SEMIMAJOR_AXIS"`
	PERIOD              string   `json:"PERIOD"`
	APOAPSIS            string   `json:"APOAPSIS"`
	PERIAPSIS           string   `json:"PERIAPSIS"`
	OBJECT_TYPE         string   `json:"OBJECT_TYPE"`
	RCS_SIZE            string   `json:"RCS_SIZE"`
	COUNTRY_CODE        string   `json:"COUNTRY_CODE"`
	LAUNCH_DATE         string   `json:"LAUNCH_DATE"`
	SITE                string   `json:"SITE"`
	DECAY_DATE          string   `json:"DECAY_DATE"`
	FILE                string   `json:"FILE"`
	GP_ID               string   `json:"GP_ID"`
	TLE_LINE0           string   `json:"TLE_LINE0"`
	TLE_LINE1           string   `json:"TLE_LINE1"`
	TLE_LINE2           string   `json:"TLE_LINE2"`
	Group               []string `json:"Group"`
	Sat                 satellite.Satellite
}
