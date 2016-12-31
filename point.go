package gopx

import "math"

// EarthRadiasKm radius of the earth in kilometers
const EarthRadiasKm = 6371

// Distance calculate the distance between to point in kilometers
func (source *Point) Distance(dest Point) float64 {
	dLat := (dest.Lat - source.Lat) * (math.Pi / 180.0)
	dLon := (dest.Lon - source.Lon) * (math.Pi / 180.0)

	lat1 := source.Lat * (math.Pi / 180.0)
	lat2 := dest.Lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadiasKm * c
}
