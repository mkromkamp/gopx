package gopx

import (
	"errors"
	"fmt"
)

// SplitAfterKM split a track into parts of n KM
func (gpx *Gpx) SplitAfterKM(km float64) ([]*Gpx, error) {
	if km <= 0 {
		return nil, errors.New("km is less than or zero")
	}

	points := gpx.GetPoints()
	if len(points) == 0 {
		return nil, errors.New("No points found. Unable to split track")
	}

	var newGpxs []*Gpx

	var gpxPoints Points
	prev := points[0]
	gpxDist := 0.0
	gpxPoints = append(gpxPoints, prev)
	for _, point := range points[1:] {
		gpxPoints = append(gpxPoints, point)
		gpxDist = gpxDist + prev.Distance(point)
		prev = point

		if gpxDist >= km {
			name := fmt.Sprintf("%s-%d", gpx.GetName(), len(newGpxs)+1)
			newGpxs = append(newGpxs, gpxPoints.createGpx(name))

			gpxPoints = gpxPoints[len(gpxPoints)-1:]
			gpxDist = 0.0
		}
	}

	// Remaining points
	newGpxs = append(newGpxs, gpxPoints.createGpx(fmt.Sprintf("%s-%d", gpx.GetName(), len(newGpxs)+1)))

	return newGpxs, nil
}

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
		name := fmt.Sprintf("%s-%d", gpx.GetName(), i)
		newGpxs = append(newGpxs, splittedPoints[i].createGpx(name))
	}

	return newGpxs, nil
}

// Split a slice of Points into even parts
func (points Points) split(parts int) []Points {
	var splittedPoints []Points
	partSize := (len(points) / parts)

	for i := 0; i < parts; i++ {
		begin := i * partSize
		end := (i + 1) * partSize

		// Last part should only contain the remaining points
		if i == partSize {
			end = len(points)
		}

		splittedPoints = append(splittedPoints, points[begin:end])
	}

	return splittedPoints
}
