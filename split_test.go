package gopx

import "testing"

func TestSplit(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 13

	gpxs, _ := gpx.Split(parts)

	if len(gpxs) != parts {
		t.Errorf("Expected %d parts but got %d parts", parts, len(gpxs))
	}
}

func TestSplitPartsZeroOrLess(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 0
	expectedError := "Parts is less than or zero"

	_, err := gpx.Split(parts)

	if err.Error() != expectedError {
		t.Errorf("Expected error: %s", expectedError)
	}
}

func TestSplitNoPoints(t *testing.T) {
	gpx, _ := ParseFile("sample/sample_empty_waypoints.gpx")
	parts := 2
	expectedError := "No points found. Unable to split track"

	_, err := gpx.Split(parts)

	if err.Error() != expectedError {
		t.Errorf("Expected error: %s", expectedError)
	}
}

func TestSplitLastPoint(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 13
	expectedPoint := gpx.Tracks[0].Segments[0].Points[len(gpx.Tracks[0].Segments[0].Points)-1]

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
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 13
	expectedPoint := gpx.Tracks[0].Segments[0].Points[0]

	gpxs, _ := gpx.Split(parts)

	// First waypoint from first gpx
	actualPoint := gpxs[0].Tracks[0].Segments[0].Points[0]

	if expectedPoint.Timestamp != actualPoint.Timestamp {
		t.Errorf("Expected point %v but got %v", expectedPoint, actualPoint)
	}
}
