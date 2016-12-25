package main

import "testing"

func TestSplit(t *testing.T) {
	gpx, _ := ParseFile("sample/sample.gpx")

	gpx.Split(2)
}
