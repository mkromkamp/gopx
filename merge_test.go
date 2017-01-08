package gopx

import (
	"reflect"
	"testing"
)

func TestMergeByDistanceOrdered(t *testing.T) {
	var points Points
	startPoint := Point{Lat: 1.0, Lon: 1.0}
	points = append(points, Point{Lat: 1.0, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.3, Lon: 1.0})

	gpxs, _ := points.createGpx("gpx").Split(2)

	merged, _ := MergeByDistance(gpxs, startPoint)
	pointsActual := merged.GetPoints()

	// What is not matching
	for i, pointExpected := range points {
		if pointsActual[i].Lat != pointExpected.Lat && pointsActual[i].Lon != pointExpected.Lon {
			t.Errorf("Expected: %v got: %v", pointExpected, pointsActual[i])
		}
	}
}

func TestMergeByDistanceUnordered(t *testing.T) {
	var points Points
	startPoint := Point{Lat: 1.0, Lon: 1.0}
	points = append(points, Point{Lat: 1.3, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.0, Lon: 1.0})

	var pointsOrdered Points
	pointsOrdered = append(pointsOrdered, Point{Lat: 1.0, Lon: 1.0})
	pointsOrdered = append(pointsOrdered, Point{Lat: 1.1, Lon: 1.0})
	pointsOrdered = append(pointsOrdered, Point{Lat: 1.2, Lon: 1.0})
	pointsOrdered = append(pointsOrdered, Point{Lat: 1.3, Lon: 1.0})

	gpxs, _ := points.createGpx("gpx").Split(2)

	merged, _ := MergeByDistance(gpxs, startPoint)
	pointsActual := merged.GetPoints()

	// What is not matching
	for i, pointExpected := range pointsOrdered {
		if pointsActual[i].Lat != pointExpected.Lat && pointsActual[i].Lon != pointExpected.Lon {
			t.Errorf("Expected: %v got: %v", pointExpected, pointsActual[i])
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

func TestMergeByTimestamp(t *testing.T) {
	var points Points
	points = append(points, Point{Timestamp: "2017-01-01T12:00:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:00:01Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:01:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T13:00:00Z"})

	gpxs, _ := points.createGpx("gpx").Split(2)

	merged, _ := MergeByTimestamp(gpxs)
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

func TestMergeByTimestampNoGpxs(t *testing.T) {
	var gpxs []*Gpx

	_, err := MergeByTimestamp(gpxs)

	if err.Error() != "No gpxs to merge" {
		t.Errorf("Expected %s but got %s", "No gpxs to merge", err.Error())
	}
}

func TestMergeByTimestampNoPoints(t *testing.T) {
	var gpxs []*Gpx
	gpxs = append(gpxs, NewGpx())
	gpxs = append(gpxs, NewGpx())

	_, err := MergeByTimestamp(gpxs)

	if err.Error() != "No Points found to merge" {
		t.Errorf("Expected %s but got %s", "No Points found to merge", err.Error())
	}
}

func TestMergeByTimestampInvalidTimestamp(t *testing.T) {
	var points Points
	points = append(points, Point{Timestamp: "2017-01-01T12:00:00Z"})
	points = append(points, Point{Timestamp: "2017-01-01T12:00:01Z"})
	points = append(points, Point{Timestamp: ""})
	points = append(points, Point{Timestamp: "2017-01-01T13:00:00Z"})

	gpxs, _ := points.createGpx("gpx").Split(2)

	_, errActual := MergeByTimestamp(gpxs)

	errExpected := "parsing time \"\" as \"2006-01-02T15:04:05.999999999Z07:00\": cannot parse \"\" as \"2006\""
	if errActual.Error() != errExpected {
		t.Errorf("Expected error: %s but got: %s", errExpected, errActual)
	}
}
