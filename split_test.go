package gopx

import "testing"

func TestSplit(t *testing.T) {
	var points Points
	points = append(points, Point{Lat: 1.0, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.3, Lon: 1.0})

	gpx := points.createGpx("gpx")
	parts := 3

	gpxs, _ := gpx.Split(parts)

	if len(gpxs) != parts {
		t.Errorf("Expected %d parts but got %d parts", parts, len(gpxs))
	}
}

func TestSplitPartsZeroOrLess(t *testing.T) {
	gpx := NewGpx()
	parts := 0
	expectedError := "Parts is less than or zero"

	_, err := gpx.Split(parts)

	if err.Error() != expectedError {
		t.Errorf("Expected error: %s", expectedError)
	}
}

func TestSplitNoPoints(t *testing.T) {
	gpx := NewGpx()
	parts := 2
	expectedError := "No points found. Unable to split track"

	_, err := gpx.Split(parts)

	if err.Error() != expectedError {
		t.Errorf("Expected error: %s", expectedError)
	}
}

func TestSplitLastPoint(t *testing.T) {
	var points Points
	points = append(points, Point{Lat: 1.0, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.3, Lon: 1.0})

	gpx := points.createGpx("gpx")
	parts := 2
	expectedPoint := points[len(points)-1]

	gpxs, _ := gpx.Split(parts)

	// Last waypoint from last gpx
	track := gpxs[len(gpxs)-1].Tracks[0]
	segment := track.Segments[len(track.Segments)-1]
	actualPoint := segment.Points[len(segment.Points)-1]

	if expectedPoint.Timestamp != actualPoint.Timestamp {
		t.Errorf("Expected point %v but got %v", expectedPoint, actualPoint)
	}
}

func TestSplitFirstWaypoint(t *testing.T) {
	var points Points
	points = append(points, Point{Lat: 1.0, Lon: 1.0})
	points = append(points, Point{Lat: 1.1, Lon: 1.0})
	points = append(points, Point{Lat: 1.2, Lon: 1.0})
	points = append(points, Point{Lat: 1.3, Lon: 1.0})

	gpx := points.createGpx("gpx")
	parts := 2
	expectedPoint := points[0]

	gpxs, _ := gpx.Split(parts)

	// First waypoint from first gpx
	actualPoint := gpxs[0].Tracks[0].Segments[0].Points[0]

	if expectedPoint.Timestamp != actualPoint.Timestamp {
		t.Errorf("Expected point %v but got %v", expectedPoint, actualPoint)
	}
}
