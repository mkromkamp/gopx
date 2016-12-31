package gopx

import "math"

// EarthRadiasKm radius of the earth in kilometers
const EarthRadiasKm = 6371

// Distance calculate the distance between to point in kilometers
// http://www.movable-type.co.uk/scripts/latlong.html
func (source *Point) Distance(dest Point) float64 {
	dLat := radians(dest.Lat - source.Lat)
	dLon := radians(dest.Lon - source.Lon)

	lat1 := radians(source.Lat)
	lat2 := radians(dest.Lat)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadiasKm * c
}

// radians converts degrees to radians.
func radians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
