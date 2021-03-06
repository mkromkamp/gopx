package gopx

import "time"

type sortByTimestamp Points

func (p sortByTimestamp) Len() int {
	return len(p)
}

func (p sortByTimestamp) Less(i, j int) bool {
	t1, _ := time.Parse(time.RFC3339Nano, p[i].Timestamp)
	t2, _ := time.Parse(time.RFC3339Nano, p[j].Timestamp)

	return t1.Before(t2)
}

func (p sortByTimestamp) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Create Gpx from Points
func (points Points) createGpx(name string) *Gpx {
	newGpx := NewGpx()
	track := Track{
		Segments: []TrackSegment{
			TrackSegment{
				Points: points,
			},
		},
	}
	metaData := Metadata{
		Name: name,
	}

	newGpx.Tracks = append(newGpx.Tracks, track)
	newGpx.Metadata = &metaData

	return newGpx
}
