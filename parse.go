package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Parse parses a GPX reader and return a Gpx object.
func Parse(r io.Reader) (*Gpx, error) {
	g := NewGpx()
	d := xml.NewDecoder(r)
	// d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(g)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse gpx data: %v", err)
	}
	return g, nil
}

// ParseFile reads a GPX file and parses it.
func ParseFile(filepath string) (*Gpx, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}

// NewGpx creates and returns a new Gpx objects.
func NewGpx() *Gpx {
	gpx := new(Gpx)
	gpx.XMLNs = "http://www.topografix.com/GPX/1/1"
	gpx.XMLNsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	gpx.XMLSchemaLoc = "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd"
	gpx.Version = "1.1"
	gpx.Creator = "https://github.com/ptrv/go-gpx"
	return gpx
}
