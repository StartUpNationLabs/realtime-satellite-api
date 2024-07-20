package service

import (
	"context"
	v1 "github.com/StartUpNationLabs/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/spacetrack"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

type SatelliteService struct {
	*v1.UnimplementedSatelliteServiceServer
	data      map[string]spacetrack.TLE
	dataArray []spacetrack.TLE
	groups    []string
}

func (api SatelliteService) GetSatelliteDetail(ctx context.Context, req *v1.SatelliteDetailRequest) (*v1.SatelliteDetail, error) {
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

func (api SatelliteService) GetSatellitePositions(_ context.Context, req *v1.GetSatellitePositionsRequest) (*v1.GetSatellitePositionsResponse, error) {
	t, filteredData := api.parseGetSatelliteParams(req)

	return &v1.GetSatellitePositionsResponse{
		Satellites: CalculateSatPositionsParallel(t, &filteredData),
	}, nil
}

func (api SatelliteService) parseGetSatelliteParams(req *v1.GetSatellitePositionsRequest) (time.Time, []spacetrack.TLE) {
	t := api.parseTime(req.Time)

	// filter the satellites by the group if provided
	filteredData := make([]spacetrack.TLE, 0)
	if len(req.Groups) > 0 {
		for _, sat := range api.dataArray {
			if compareGroups(req.Groups, sat.Group) {
				filteredData = append(filteredData, sat)
			}
		}
	} else {
		filteredData = api.dataArray
	}
	return t, filteredData
}

func (api SatelliteService) parseTime(reqTime *timestamppb.Timestamp) time.Time {
	if reqTime == nil {
		reqTime = &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		}
	}
	t := time.Unix(reqTime.Seconds, int64(reqTime.Nanos))
	log.Info("Request time: ", t)
	return t
}

func (api SatelliteService) GetSatelliteGroups(context.Context, *emptypb.Empty) (*v1.GetSatelliteGroupsResponse, error) {
	return &v1.GetSatelliteGroupsResponse{
		Groups: api.groups,
	}, nil
}

func (api SatelliteService) GetMinimalSatellites(_ context.Context, req *v1.GetSatellitePositionsRequest) (*v1.GetMinimalSatellitesResponse, error) {
	_, filteredData := api.parseGetSatelliteParams(req)

	// map over the data and return only the minimal data
	minimalData := make([]*v1.MinimalSatellite, 0)
	for _, data := range filteredData {
		minimalData = append(minimalData, &v1.MinimalSatellite{
			NoradCatId: data.NORAD_CAT_ID,
			ObjectName: data.OBJECT_NAME,
		})
	}

	return &v1.GetMinimalSatellitesResponse{
		Satellites: minimalData,
	}, nil
}

func (api SatelliteService) GetSatellitePath(_ context.Context, req *v1.SatellitePathRequest) (*v1.GetSatellitePathResponse, error) {
	// meanMotion is the revolution per day of the satellite
	meanMotion, err := strconv.ParseFloat(api.data[req.Id].MEAN_MOTION, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error parsing mean motion: %s", err)
	}

	// period is the time it takes for the satellite to complete one orbit
	period := (1.0 / meanMotion) * 24 * 60 * 60

	// get the current time
	t := api.parseTime(req.GetTime())
	startTime := t.Add(-time.Duration(period) * time.Second)
	endTime := t.Add(time.Duration(period) * time.Second)

	// calculate the positions of the satellite
	positions := make([]*v1.GeoPoint, 0)

	resolution := req.Resolution
	log.Info("Resolution: ", resolution)
	if resolution == 0 {
		resolution = 60
	}

	step := period / float64(resolution)
	total := 0
	for i := startTime.Unix(); i < endTime.Unix(); i += int64(step) {
		total++
		t := time.Unix(i, 0)
		tle := api.data[req.Id]
		calculatedSat := CalculateSatPosition(t, &tle)
		positions = append(positions, &v1.GeoPoint{
			Lat:      calculatedSat.Lat,
			Lon:      calculatedSat.Lon,
			Altitude: calculatedSat.Altitude,
			Velocity: calculatedSat.Velocity,
		})
	}
	log.Info("Total positions: ", total)

	return &v1.GetSatellitePathResponse{
		Path: positions,
	}, nil

}
