package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/joshuaferrara/go-satellite"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/celestrack"
	"github.com/tsukoyachi/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"time"
)

type SatelliteService struct {
	*v1.UnimplementedSatelliteServiceServer
	data map[string]spacetrack.TLE
}

func NewSatelliteService() *SatelliteService {
	client := resty.New()

	credentials, err := spacetrack.Login(client, spacetrack.Credentials{
		Identity: os.Getenv("identity"),
		Password: os.Getenv("password"),
	})

	if err != nil {
		log.Error(err)

	}
	data, err := spacetrack.FetchData(client, "data.json", credentials)
	if err != nil {
		log.Error(err)
	}

	calculatedMap := make(map[string]spacetrack.TLE)
	for _, tle := range data {
		if tle.NORAD_CAT_ID == "" {
			log.Fatalf("NORAD_CAT_ID is empty in %v", tle)
		}
		// check if the object id is already in the map
		if _, ok := calculatedMap[tle.NORAD_CAT_ID]; ok {
			log.Fatalf("NORAD_CAT_ID %v is already in the map", tle.OBJECT_ID)
		}
		calculatedMap[tle.NORAD_CAT_ID] = tle

	}

	// fetch from celestrak
	celestrakData, err := celestrack.Scrap()
	if err != nil {
		log.Error(err)
	}
	for k, v := range celestrakData {
		sat, ok := calculatedMap[k]
		if !ok {
			calculatedMap[k] = v
			log.Info("Added ", k)
			continue
		}
		sat.Group = append(sat.Group, v.Group...)
		calculatedMap[k] = sat
	}
	log.Info("Total satellites: ", len(calculatedMap))
	return &SatelliteService{
		data: calculatedMap,
	}
}

func (api SatelliteService) GetSatelliteDetail(ctx context.Context, req *v1.Satellite) (*v1.SatelliteDetail, error) {
	// Check if the satellite id is in the map
	tle, ok := api.data[req.Id]
	if !ok {
		// If not, return an error
		return nil, status.Errorf(codes.NotFound, "Satellite with ID %s not found", req.Id)
	}

	return &v1.SatelliteDetail{
		CcsdsOmmVers:       tle.CCSDS_OMM_VERS,
		Comment:            tle.COMMENT,
		CreationDate:       tle.CREATION_DATE,
		Originator:         tle.ORIGINATOR,
		ObjectName:         tle.OBJECT_NAME,
		ObjectId:           tle.OBJECT_ID,
		CenterName:         tle.CENTER_NAME,
		RefFrame:           tle.REF_FRAME,
		TimeSystem:         tle.TIME_SYSTEM,
		MeanElementTheory:  tle.MEAN_ELEMENT_THEORY,
		Epoch:              tle.EPOCH,
		MeanMotion:         tle.MEAN_MOTION,
		Eccentricity:       tle.ECCENTRICITY,
		Inclination:        tle.INCLINATION,
		RaOfAscNode:        tle.RA_OF_ASC_NODE,
		ArgOfPericenter:    tle.ARG_OF_PERICENTER,
		MeanAnomaly:        tle.MEAN_ANOMALY,
		EphemerisType:      tle.EPHEMERIS_TYPE,
		ClassificationType: tle.CLASSIFICATION_TYPE,
		NoradCatId:         tle.NORAD_CAT_ID,
		ElementSetNo:       tle.ELEMENT_SET_NO,
		RevAtEpoch:         tle.REV_AT_EPOCH,
		Bstar:              tle.BSTAR,
		MeanMotionDot:      tle.MEAN_MOTION_DOT,
		MeanMotionDdot:     tle.MEAN_MOTION_DDOT,
		SemimajorAxis:      tle.SEMIMAJOR_AXIS,
		Period:             tle.PERIOD,
		Apoapsis:           tle.APOAPSIS,
		Periapsis:          tle.PERIAPSIS,
		ObjectType:         tle.OBJECT_TYPE,
		RcsSize:            tle.RCS_SIZE,
		CountryCode:        tle.COUNTRY_CODE,
		LaunchDate:         tle.LAUNCH_DATE,
		Site:               tle.SITE,
		DecayDate:          tle.DECAY_DATE,
		File:               tle.FILE,
		GpId:               tle.GP_ID,
		TleLine0:           tle.TLE_LINE0,
		TleLine1:           tle.TLE_LINE1,
		TleLine2:           tle.TLE_LINE2,
		Groups:             tle.Group,
	}, nil
}

func (api SatelliteService) GetSatellitePositions(ctx context.Context, req *v1.GetSatellitePositionsRequest) (*v1.GetSatellitePositionsResponse, error) {
	calculatedPositions := make([]*v1.Satellite, 0)
	reqTime := req.GetTime()
	if reqTime == nil {
		reqTime = &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		}
	}
	// iterate over api.data
	for id, data := range api.data {
		log.Info("Calculating position for satellite", id)
		log.Info(data.TLE_LINE1, data.TLE_LINE2)
		goSat := satellite.TLEToSat(data.TLE_LINE1, data.TLE_LINE2, satellite.GravityWGS84)
		// propagate the satellite position
		t := time.Unix(reqTime.Seconds, int64(reqTime.Nanos))

		log.Info(goSat.ErrorStr, goSat.Error)
		log.Info(t)
		pos, vec := satellite.Propagate(goSat, t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
		log.Info(pos, vec)
		gst := satellite.GSTimeFromDate(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

		// convert Earth Centered Inertial coordinates to Lat/Long
		altitude, velocity, ret := satellite.ECIToLLA(pos, gst)
		calculatedPositions = append(calculatedPositions, &v1.Satellite{
			Id:       id,
			Lat:      ret.Latitude,
			Lon:      ret.Longitude,
			Altitude: altitude,
			Velocity: velocity,
		})
	}
	return &v1.GetSatellitePositionsResponse{
		Satellites: calculatedPositions,
	}, nil
}
