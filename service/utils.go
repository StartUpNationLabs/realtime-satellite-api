package service

import (
	"fmt"
	v1 "github.com/StartUpNationLabs/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/spacetrack"
	"github.com/joshuaferrara/go-satellite"
	"runtime"
	"slices"
	"time"
)

func CalculateSatPosition(t time.Time, sat *spacetrack.TLE) *v1.Satellite {
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

func CalculateChunk(t time.Time, sats *[]spacetrack.TLE) []*v1.Satellite {
	calculatedPositions := make([]*v1.Satellite, 0)
	for _, data := range *sats {
		calculatedPositions = append(calculatedPositions, CalculateSatPosition(t, &data))
	}
	return calculatedPositions
}

func Chunkify(sats *[]spacetrack.TLE, numberOfChunks int) [][]spacetrack.TLE {
	chunkSize := len(*sats) / numberOfChunks
	chunks := make([][]spacetrack.TLE, 0)
	for i := 0; i < len(*sats); i += chunkSize {
		end := i + chunkSize
		if end > len(*sats) {
			end = len(*sats)
		}
		chunks = append(chunks, (*sats)[i:end])
	}
	return chunks
}

type ChunkResult struct {
	Index int
	Data  []*v1.Satellite
}

func CalculateSatPositionsParallel(t time.Time, data *[]spacetrack.TLE) []*v1.Satellite {
	// chunkify the data
	time1 := time.Now()
	// chunk to the number of CPUs
	chunks := Chunkify(data, runtime.NumCPU())
	time2 := time.Now()
	fmt.Println("Time to chunkify: ", time2.Sub(time1))

	// make the channels
	ch := make(chan ChunkResult)
	for i, chunk := range chunks {
		go func(index int, chunk []spacetrack.TLE) {
			ch <- ChunkResult{Index: index, Data: CalculateChunk(t, &chunk)}
		}(i, chunk)
	}
	time3 := time.Now()

	results := make([][]*v1.Satellite, len(chunks))
	for range chunks {
		result := <-ch
		results[result.Index] = result.Data
	}

	calculatedPositions := make([]*v1.Satellite, 0)
	for _, result := range results {
		calculatedPositions = append(calculatedPositions, result...)
	}

	time4 := time.Now()
	fmt.Println("Time to calculate chunks: ", time4.Sub(time3))
	fmt.Println("Total time Parrael: ", time4.Sub(time1))
	return calculatedPositions
}

func CalculateSatPositionsLoop(t time.Time, data *[]spacetrack.TLE) []*v1.Satellite {
	time1 := time.Now()
	calculatedPositions := make([]*v1.Satellite, 0)
	for _, data := range *data {
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

func compareGroups(a []string, b []string) bool {
	// if every of a is not in b, return false
	for _, group := range a {
		if !slices.Contains(b, group) {
			return false
		}
	}
	return true
}
