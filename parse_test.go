package gopx

import "testing"

func TestParse(t *testing.T) {
	var err error
	var g *Gpx
	g, err = ParseFile("sample/sample.gpx")

	if err != nil {
		t.Error("Error parsing GPX file: ", err)
	}

	if g == nil {
		t.Error("Expected new Gpx")
	}
}

func TestParseMetadata(t *testing.T) {
	var g *Gpx
	g, _ = ParseFile("sample/sample.gpx")

	timestampA := g.Metadata.Timestamp
	timestampE := "2016-12-23T15:12:25Z"
	if timestampA != timestampE {
		t.Errorf("timestamp expected: %s, actual: %s", timestampE, timestampA)
	}
}

func TestParseTracks(t *testing.T) {
	var g *Gpx
	g, _ = ParseFile("sample/sample.gpx")

	tracknameA := g.Tracks[0].Name
	tracknameE := "Werk-woon"
	if tracknameA != tracknameE {
		t.Errorf("Trackname expected: %s, actual: %s", tracknameE, tracknameA)
	}
}

func TestParsePoints(t *testing.T) {
	var g *Gpx
	g, _ = ParseFile("sample/sample.gpx")

	wayEleA := g.Tracks[0].Segments[0].Points[0].Elevation
	wayEleE := -2.3
	if wayEleE != wayEleA {
		t.Errorf("Number of tracks expected: %f, actual: %f", wayEleE, wayEleA)
	}
}

func TestParseInvalidFile(t *testing.T) {
	var errActual error
	_, errActual = ParseFile("sample/does_not_exists.gpx")

	errExpected := "open sample/does_not_exists.gpx: no such file or directory"
	if errActual.Error() != errExpected {
		t.Errorf("Expected error: %s but got: %s", errExpected, errActual)
	}
}
