package main

import "testing"

func TestParse(t *testing.T) {
	var err error
	var g *Gpx
	g, err = ParseFile("sample/sample.gpx")

	if err != nil {
		t.Error("Error parsing GPX file: ", err)
	}

	timestampA := g.Metadata.Timestamp
	timestampE := "2016-12-23T15:12:25Z"
	if timestampA != timestampE {
		t.Errorf("timestamp expected: %s, actual: %s", timestampE, timestampA)
	}

	trknameA := g.Tracks[0].Name
	trknameE := "Werk-woon"
	if trknameA != trknameE {
		t.Errorf("Trackname expected: %s, actual: %s", trknameE, trknameA)
	}

	numPointsA := g.Tracks[0].Segments[0].Waypoints[0].Elevation
	numPointsE := -2.3
	if numPointsE != numPointsA {
		t.Errorf("Number of tracks expected: %f, actual: %f", numPointsE, numPointsA)
	}
}
