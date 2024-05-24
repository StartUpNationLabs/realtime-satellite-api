package service

import (
	"context"
	"fmt"
	v1 "github.com/StartUpNationLabs/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/spacetrack"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"runtime"
	"time"
)

type SatelliteService struct {
	*v1.UnimplementedSatelliteServiceServer
	data map[string]spacetrack.TLE
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

func (api SatelliteService) GetSatellitePositions(_ context.Context, req *v1.GetSatellitePositionsRequest) (*v1.GetSatellitePositionsResponse, error) {
	reqTime := req.GetTime()
	if reqTime == nil {
		reqTime = &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		}
	}
	t := time.Unix(reqTime.Seconds, int64(reqTime.Nanos))

	return &v1.GetSatellitePositionsResponse{
		Satellites: CalculateSatPositionsParrallel(t, api.data),
	}, nil
}

func CalculateSatPosition(t time.Time, sat spacetrack.TLE) *v1.Satellite {
	// propagate the satellite position
	pos, _ := satellite.Propagate(sat.Sat, t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	gst := satellite.GSTimeFromDate(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

	// convert Earth Centered Inertial coordinates to Lat/Long
	altitude, velocity, ret := satellite.ECIToLLA(pos, gst)
	return &v1.Satellite{
		Id:       sat.NORAD_CAT_ID,
		Name:     sat.OBJECT_NAME,
		Lat:      ret.Latitude,
		Lon:      ret.Longitude,
		Altitude: altitude,
		Velocity: velocity,
	}
}

func CalculateChunk(t time.Time, sats []spacetrack.TLE) []*v1.Satellite {
	calculatedPositions := make([]*v1.Satellite, 0)
	for _, data := range sats {
		calculatedPositions = append(calculatedPositions, CalculateSatPosition(t, data))
	}
	return calculatedPositions
}

func Chunkify(t time.Time, sats []spacetrack.TLE, numberOfChunks int) [][]spacetrack.TLE {
	chunkSize := len(sats) / numberOfChunks
	chunks := make([][]spacetrack.TLE, 0)
	for i := 0; i < len(sats); i += chunkSize {
		end := i + chunkSize
		if end > len(sats) {
			end = len(sats)
		}
		chunks = append(chunks, sats[i:end])
	}
	return chunks
}

func CalculateSatPositionsParrallel(t time.Time, data map[string]spacetrack.TLE) []*v1.Satellite {
	// chunkify the data
	time1 := time.Now()
	// chunk to the number of CPUs
	chunks := Chunkify(t, maps.Values(data), runtime.NumCPU())
	time2 := time.Now()
	fmt.Println("Time to chunkify: ", time2.Sub(time1))

	// make the channels
	ch := make(chan []*v1.Satellite)
	for _, chunk := range chunks {
		go func(chunk []spacetrack.TLE) {
			ch <- CalculateChunk(t, chunk)
		}(chunk)
	}
	time3 := time.Now()

	calculatedPositions := make([]*v1.Satellite, 0)
	for range chunks {
		calculatedPositions = append(calculatedPositions, <-ch...)

	}
	time4 := time.Now()
	fmt.Println("Time to calculate chunks: ", time4.Sub(time3))
	fmt.Println("Total time Parrael: ", time4.Sub(time1))
	return calculatedPositions

}

func CalculateSatPositionsLoop(t time.Time, data map[string]spacetrack.TLE) []*v1.Satellite {
	time1 := time.Now()
	calculatedPositions := make([]*v1.Satellite, 0)
	for _, data := range data {
		// propagate the satellite position

		pos, _ := satellite.Propagate(data.Sat, t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
		gst := satellite.GSTimeFromDate(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

		// convert Earth Centered Inertial coordinates to Lat/Long
		altitude, velocity, ret := satellite.ECIToLLA(pos, gst)
		calculatedPositions = append(calculatedPositions, &v1.Satellite{
			Id:       data.NORAD_CAT_ID,
			Name:     data.OBJECT_NAME,
			Lat:      ret.Latitude,
			Lon:      ret.Longitude,
			Altitude: altitude,
			Velocity: velocity,
		})
	}
	time2 := time.Now()
	fmt.Println("Time to calculate positions Loop: ", time2.Sub(time1))
	return calculatedPositions
}
