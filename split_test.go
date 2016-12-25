package main

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

func TestSplitNoWaypoints(t *testing.T) {
	gpx, _ := ParseFile("sample/sample_empty_waypoints.gpx")
	parts := 2
	expectedError := "No waypoints found. Unable to split track"

	_, err := gpx.Split(parts)

	if err.Error() != expectedError {
		t.Errorf("Expected error: %s", expectedError)
	}
}

func TestSplitLastWaypoint(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 13
	expectedWaypoint := gpx.Tracks[0].Segments[0].Waypoints[len(gpx.Tracks[0].Segments[0].Waypoints)-1]

	gpxs, _ := gpx.Split(parts)

	// Last waypoint from last gpx
	track := gpxs[len(gpxs)-1].Tracks[0]
	segment := track.Segments[len(track.Segments)-1]
	actualWaypoint := segment.Waypoints[len(segment.Waypoints)-1]

	if expectedWaypoint.Timestamp != actualWaypoint.Timestamp {
		t.Errorf("Expected waypoint %v but got %v", expectedWaypoint, actualWaypoint)
	}
}

func TestSplitFirstWaypoint(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")
	parts := 13
	expectedWaypoint := gpx.Tracks[0].Segments[0].Waypoints[0]

	gpxs, _ := gpx.Split(parts)

	// First waypoint from first gpx
	actualWaypoint := gpxs[0].Tracks[0].Segments[0].Waypoints[0]

	if expectedWaypoint.Timestamp != actualWaypoint.Timestamp {
		t.Errorf("Expected waypoint %v but got %v", expectedWaypoint, actualWaypoint)
	}
}
