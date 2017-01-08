package gopx

import (
	"errors"
	"sort"
	"time"
)

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	t1, _ := time.Parse(time.RFC3339Nano, p[i].Timestamp)
	t2, _ := time.Parse(time.RFC3339Nano, p[j].Timestamp)

	return t1.Before(t2)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// MergeByTimestamp merges gpxs based on there timetamp
func MergeByTimestamp(gpxs []*Gpx) (*Gpx, error) {
	points, err := getPoints(gpxs)
	if err != nil {
		return nil, err
	}

	for _, point := range points {
		_, err := time.Parse(time.RFC3339Nano, point.Timestamp)
		if err != nil {
			return nil, err
		}
	}

	sort.Sort(points)

	return points.createGpx("merged"), nil
}

// MergeByDistance merges gpxs based on distance
func MergeByDistance(gpxs []*Gpx, start Point) (*Gpx, error) {
	points, err := getPoints(gpxs)
	if err != nil {
		return nil, err
	}

	currentPoint := start
	var sorted Points
	for len(points) > 0 {
		closest := points[0]
		closestIndex := 0
		closestDist := currentPoint.Distance(closest)

		// Find the next closest point
		for i, p := range points {
			pDist := currentPoint.Distance(p)
			if pDist < closestDist {
				closest = p
				closestIndex = i
				closestDist = pDist
			}
		}

		points = append(points[:closestIndex], points[closestIndex+1:]...)
		sorted = append(sorted, closest)
		currentPoint = closest
	}

	return sorted.createGpx("merged"), nil
}

// getPoints get points from gpxs.
func getPoints(gpxs []*Gpx) (Points, error) {
	if len(gpxs) == 0 {
		return nil, errors.New("No gpxs to merge")
	}

	var points Points
	for _, gpx := range gpxs {
		points = append(points, gpx.GetPoints()...)
	}

	if len(points) == 0 {
		return nil, errors.New("No Points found to merge")
	}

	return points, nil
}
