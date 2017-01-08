package gopx

import (
	"sort"
	"testing"
)

func TestSortByTimestampOrdered(t *testing.T) {
	var points Points
	points = append(points, Point{Timestamp: "2017-01-01T12:00:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:00:01Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:01:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T13:00:00Z"})

	sorted := points
	sort.Sort(sortByTimestamp(sorted))

	for i, p := range sorted {
		if points[i].Timestamp != p.Timestamp {
			t.Errorf("Expected point: %v but got: %v", points[i], p)
		}
	}
}

func TestSortByTimestampUnordered(t *testing.T) {
	var points Points
	points = append(points, Point{Timestamp: "2017-01-01T12:00:01Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:00:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T13:00:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:01:00Z"})

	var pointsExpected Points
	pointsExpected = append(pointsExpected, Point{Timestamp: "2017-01-01T12:00:00Z"})
	pointsExpected = append(pointsExpected, Point{Timestamp: "2017-01-01T12:00:01Z"})
	pointsExpected = append(pointsExpected, Point{Timestamp: "2017-01-01T12:01:00Z"})
	pointsExpected = append(pointsExpected, Point{Timestamp: "2017-01-01T13:00:00Z"})

	sorted := points
	sort.Sort(sortByTimestamp(sorted))

	for i, p := range sorted {
		if pointsExpected[i].Timestamp != p.Timestamp {
			t.Errorf("Expected point: %v but got: %v", pointsExpected[i], p)
		}
	}
}
