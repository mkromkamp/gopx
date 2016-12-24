package main

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
		t.Error("Expected %s XMLNs but got %s", xmlnsE, xmlnsA)
	}
}

func TestNewGpxXMLNsXsi(t *testing.T) {
	gpx := NewGpx()

	xmlnsXsiE := "http://www.w3.org/2001/XMLSchema-instance"
	xmlnsXsiA := gpx.XMLNsXsi
	if xmlnsXsiE != xmlnsXsiA {
		t.Error("Expected %s XMLNsXsi but got %s", xmlnsXsiE, xmlnsXsiA)
	}
}

func TestNewGpxXMLSchemaLoc(t *testing.T) {
	gpx := NewGpx()

	xmlSchemaLocE := "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd"
	xmlSchemaLocA := gpx.XMLSchemaLoc
	if xmlSchemaLocE != xmlSchemaLocA {
		t.Error("Expected %s XMLSchemaLoc but got %s", xmlSchemaLocE, xmlSchemaLocA)
	}
}

func TestNewGpxVersion(t *testing.T) {
	gpx := NewGpx()

	versionE := "1.1"
	versionA := gpx.Version
	if versionE != versionA {
		t.Error("Expected %s Version but got %s", versionE, versionA)
	}
}
