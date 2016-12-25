package main

import (
	"errors"
	"fmt"
)

// Split a Gpx track into even parts
func (gpx *Gpx) Split(parts int) ([]*Gpx, error) {
	if parts <= 0 {
		return nil, errors.New("Parts is less than or zero")
	}

	waypoints := gpx.getWaypoints()
	if len(waypoints) == 0 {
		return nil, errors.New("No waypoints found. Unable to split track")
	}

	splittedWaypoints := waypoints.split(parts)

	var newGpxs []*Gpx
	for i := 0; i < parts; i++ {
		newGpx := NewGpx()
		track := Track{
			Name: fmt.Sprintf("%s-%d", gpx.Tracks[0].Name, i),
			Segments: []TrackSegment{
				TrackSegment{
					Waypoints: splittedWaypoints[1],
				},
			},
		}

		newGpx.Tracks = append(newGpx.Tracks, track)
		newGpxs = append(newGpxs, newGpx)
	}

	return newGpxs, nil
}

// Get all the waypoints for a Gpx
func (gpx *Gpx) getWaypoints() Waypoints {
	var waypoints Waypoints
	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			waypoints = append(waypoints, segment.Waypoints...)
		}
	}

	return waypoints
}

// Split a slice of Waypoints into even parts
func (waypoints Waypoints) split(parts int) []Waypoints {
	var splittedWaypoints []Waypoints
	partSize := len(waypoints) / parts

	for i := 0; i < parts; i++ {
		begin := i * partSize
		end := (i + 1) * partSize

		splittedWaypoints = append(splittedWaypoints, waypoints[begin:end])
	}

	return splittedWaypoints
}
