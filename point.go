package gopx

import "math"

// EarthRadiasKm radius of the earth in kilometers
const EarthRadiasKm = float64(6371)

// Distance calculate the distance between to point in kilometers
// http://www.movable-type.co.uk/scripts/latlong.html
func (source *Point) Distance(dest Point) float64 {
	deltaLat := radians(dest.Lat - source.Lat)
	deltaLon := radians(dest.Lon - source.Lon)

	latSource := radians(source.Lat)
	latDest := radians(dest.Lat)

	a1 := math.Sin(deltaLat/2) * math.Sin(deltaLat/2)
	a2 := math.Sin(deltaLon/2) * math.Sin(deltaLon/2) * math.Cos(latSource) * math.Cos(latDest)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadiasKm * c
}

// radians converts degrees to radians.
func radians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
