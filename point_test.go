package gopx

import "testing"

func TestDistance(t *testing.T) {
	source := Point{Lat: 37.4219999, Lon: -122.0840575}
	dest := Point{Lat: 37.7888305, Lon: -122.426465}
	expectedDistance := 50.730879

	distance := source.Distance(dest)

	// Account for a slight offset, this is ok
	if !(distance < (expectedDistance+0.1) && distance > (expectedDistance-0.1)) {
		t.Errorf("Expected a distance of %f but got %f", expectedDistance, distance)
	}
}
