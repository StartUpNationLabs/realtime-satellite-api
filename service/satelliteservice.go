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
			log.Fatalf("OBJECT_ID is empty in %v", tle)
		}
		// check if the object id is already in the map
		if _, ok := calculatedMap[tle.NORAD_CAT_ID]; ok {
			log.Fatalf("OBJECT_ID %v is already in the map", tle.OBJECT_ID)
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

	// If it is, return the satellite detail
	return &v1.SatelliteDetail{
		Id:          tle.NORAD_CAT_ID,
		Name:        tle.OBJECT_NAME,
		Description: tle.COMMENT,
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
