package gopx

import (
	"errors"
	"fmt"
)

// Split a Gpx track into even parts
func (gpx *Gpx) Split(parts int) ([]*Gpx, error) {
	if parts <= 0 {
		return nil, errors.New("Parts is less than or zero")
	}

	points := gpx.GetPoints()
	if len(points) == 0 {
		return nil, errors.New("No points found. Unable to split track")
	}

	splittedPoints := points.split(parts)

	var newGpxs []*Gpx
	for i := 0; i < parts; i++ {
		name := fmt.Sprintf("%s-%d", gpx.Tracks[0].Name, i)
		newGpxs = append(newGpxs, splittedPoints[i].createGpx(name))
	}

	return newGpxs, nil
}

// Split a slice of Points into even parts
func (points Points) split(parts int) []Points {
	var splittedPoints []Points
	partSize := (len(points) / parts) + 1

	for i := 0; i < parts; i++ {
		begin := i * partSize
		end := (i + 1) * partSize

		// Last part should only contain the remaining points
		if end > len(points) {
			end = len(points)
		}

		splittedPoints = append(splittedPoints, points[begin:end])
	}

	return splittedPoints
}

// Create Gpx from Points
func (points Points) createGpx(name string) *Gpx {
	newGpx := NewGpx()
	track := Track{
		Name: name,
		Segments: []TrackSegment{
			TrackSegment{
				Points: points,
			},
		},
	}

	newGpx.Tracks = append(newGpx.Tracks, track)

	return newGpx
}
