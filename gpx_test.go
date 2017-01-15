package gopx

import "testing"

func TestNewGpx(t *testing.T) {
	gpx := NewGpx()

	if gpx == nil {
		t.Error("Expected Gpx struct.")
	}
}

func TestNewGpxXMLNs(t *testing.T) {
	gpx := NewGpx()

	xmlnsE := "http://www.topografix.com/GPX/1/1"
	xmlnsA := gpx.XMLNs
	if xmlnsE != xmlnsA {
		t.Errorf("Expected %s XMLNs but got %s", xmlnsE, xmlnsA)
	}
}

func TestNewGpxXMLNsXsi(t *testing.T) {
	gpx := NewGpx()

	xmlnsXsiE := "http://www.w3.org/2001/XMLSchema-instance"
	xmlnsXsiA := gpx.XMLNsXsi
	if xmlnsXsiE != xmlnsXsiA {
		t.Errorf("Expected %s XMLNsXsi but got %s", xmlnsXsiE, xmlnsXsiA)
	}
}

func TestNewGpxXMLSchemaLoc(t *testing.T) {
	gpx := NewGpx()

	xmlSchemaLocE := "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd"
	xmlSchemaLocA := gpx.XMLSchemaLoc
	if xmlSchemaLocE != xmlSchemaLocA {
		t.Errorf("Expected %s XMLSchemaLoc but got %s", xmlSchemaLocE, xmlSchemaLocA)
	}
}

func TestNewGpxVersion(t *testing.T) {
	gpx := NewGpx()

	versionE := "1.1"
	versionA := gpx.Version
	if versionE != versionA {
		t.Errorf("Expected %s Version but got %s", versionE, versionA)
	}
}

func TestGetPoints(t *testing.T) {
	var gpx *Gpx
	gpx, _ = ParseFile("sample/sample.gpx")
	expectedPoints := gpx.Tracks[0].Segments[0].Points

	points := gpx.GetPoints()

	if len(points) != len(expectedPoints) {
		t.Errorf("Expected %d points but got %d points", len(expectedPoints), len(points))
	}
}

func TestGetNameEmpty(t *testing.T) {
	gpx := NewGpx()

	if gpx.GetName() != "" {
		t.Errorf("Expected empty gpx name but got: %s", gpx.GetName())
	}
}

func TestGetName(t *testing.T) {
	gpx := NewGpx()
	nameExpected := "TestName"
	gpx.Metadata = &Metadata{
		Name: nameExpected,
	}

	if gpx.GetName() != nameExpected {
		t.Errorf("Expected gpx name: %s but got: %s", nameExpected, gpx.GetName())
	}
}
