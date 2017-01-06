package gopx

import (
	"reflect"
	"testing"
)

func TestMergeByDistance(t *testing.T) {
	var points Points
	startPoint := Point{Lat: 1.0, Lon: 1.0}
	points = append(points, Point{Lat: 1.0, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.3, Lon: 1.0})

	gpxs, _ := points.createGpx("gpx").Split(2)

	merged, _ := MergeByDistance(gpxs, startPoint)
	pointsActual := merged.GetPoints()

	if !reflect.DeepEqual(points, pointsActual) {
		t.Error("Points do not match")

		// What is not matching
		for i, pointExpected := range points {
			if pointsActual[i].Timestamp != pointExpected.Timestamp {
				t.Errorf("Expected: %v got: %v", pointExpected, pointsActual[i])
			}
		}
	}
}

func TestMergeByDistanceNoGpxs(t *testing.T) {
	var gpxs []*Gpx
	startPoint := Point{Lat: 1.0, Lon: 1.0}

	_, err := MergeByDistance(gpxs, startPoint)

	if err.Error() != "No gpxs to merge" {
		t.Errorf("Expected %s but got %s", "No gpxs to merge", err.Error())
	}
}

func TestMergeByDistanceNoPoints(t *testing.T) {
	var gpxs []*Gpx
	startPoint := Point{Lat: 1.0, Lon: 1.0}
	gpxs = append(gpxs, NewGpx())
	gpxs = append(gpxs, NewGpx())

	_, err := MergeByDistance(gpxs, startPoint)

	if err.Error() != "No Points found to merge" {
		t.Errorf("Expected %s but got %s", "No Points found to merge", err.Error())
	}
}
