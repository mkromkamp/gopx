package gopx

// NewGpx creates and returns a new Gpx objects.
func NewGpx() *Gpx {
	gpx := new(Gpx)
	gpx.XMLNs = "http://www.topografix.com/GPX/1/1"
	gpx.XMLNsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	gpx.XMLSchemaLoc = "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd"
	gpx.Version = "1.1"
	// gpx.Creator = "https://github.com/mkromkamp/gpx-go"
	return gpx
}

// GetPoints get all the points from this Gpx
func (gpx *Gpx) GetPoints() Points {
	var points Points
	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			points = append(points, segment.Points...)
		}
	}

	return points
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

// GetName of the gpx
func (gpx *Gpx) GetName() string {
	if gpx.Metadata != nil {
		return gpx.Metadata.Name
	}

	return ""
}
