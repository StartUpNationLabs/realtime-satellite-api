package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
)

type SatelliteService struct {
	*v1.UnimplementedSatelliteServiceServer
}

func NewSatelliteService() *SatelliteService {
	return &SatelliteService{}
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
func (SatelliteService) GetSatelliteDetail(context.Context, *v1.Satellite) (*v1.SatelliteDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatelliteDetail not implemented")
}
