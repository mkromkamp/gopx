package gopx

import (
	"encoding/xml"
	"io"
	"os"
)

// Parse parses a GPX reader and return a Gpx object.
func Parse(r io.Reader) (*Gpx, error) {
	gpx := NewGpx()
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(gpx)
	if err != nil {
		return nil, err
	}
	return gpx, nil
}

// ParseFile reads a GPX file and parses it.
func ParseFile(filepath string) (*Gpx, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Parse(file)
}
