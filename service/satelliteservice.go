package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"os"
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
	log.Println(len(data))

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
	}, nil
}

func (api SatelliteService) GetSatellitePositions(ctx context.Context, req *v1.GetSatellitePositionsRequest) (*v1.GetSatellitePositionsResponse, error) {
	// print("GetSatellitePositions")
	log.Info(req.GetTime())
	return &v1.GetSatellitePositionsResponse{
		Satellites: []*v1.Satellite{
			{
				Id:       "1",
				Lat:      0,
				Lon:      0,
				Altitude: float64(req.GetTime().GetSeconds()),
			},
		},
	}, nil
}
