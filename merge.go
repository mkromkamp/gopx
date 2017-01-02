package gopx

import "errors"

// MergeByDistance merges gpxs based on distance
func MergeByDistance(gpxs []*Gpx, start Point) (*Gpx, error) {
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
