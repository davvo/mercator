package mercator

import "math"

var tileSize = 256.0
var initialResolution = 2 * math.Pi * 6378137 / tileSize
var originShift = 2 * math.Pi * 6378137 / 2

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}

// Resolution (meters/pixel) for given zoom level (measured at Equator)
func Resolution(zoom int) float64 {
	return initialResolution / math.Pow(2, float64(zoom))
}

// Zoom level for given resolution (measured at Equator)
func Zoom(resolution float64) int {
	zoom := round(math.Log(initialResolution/resolution) / math.Log(2))
	return int(zoom)
}

// Converts given lat/lon in WGS84 Datum to XY in Spherical Mercator EPSG:900913
func LatLonToMeters(lat, lon float64) (x, y float64) {
	x = lon * originShift / 180
	y = math.Log(math.Tan((90+lat)*math.Pi/360)) / (math.Pi / 180)
	y = y * originShift / 180
	return x, y
}

// Converts XY point from Spherical Mercator EPSG:900913 to lat/lon in WGS84 Datum
func MetersToLatLon(mx, my float64) (lat, lon float64) {
	lon = (mx / originShift) * 180
	lat = (my / originShift) * 180
	lat = 180 / math.Pi * (2*math.Atan(math.Exp(lat*math.Pi/180)) - math.Pi/2)
	return lat, lon
}

// Converts pixel coordinates in given zoom level of pyramid to EPSG:900913
func PixelsToMeters(px, py float64, zoom int) (x, y float64) {
	res := Resolution(zoom)
	x = px*res - originShift
	y = py*res - originShift
	return x, y
}

// Converts EPSG:900913 to pixel coordinates in given zoom level
func MetersToPixels(x, y float64, zoom int) (px, py float64) {
	res := Resolution(zoom)
	px = (x + originShift) / res
	py = (y + originShift) / res
	return px, py
}

// Converts given lat/lon in WGS84 Datum to pixel coordinates in given zoom level
func LatLonToPixels(lat, lon float64, zoom int) (px, py float64) {
	x, y := LatLonToMeters(lat, lon)
	return MetersToPixels(x, y, zoom)
}

// Converts pixel coordinates in given zoom level to lat/lon in WGS84 Datum
func PixelsToLatLon(px, py float64, zoom int) (lat, lon float64) {
	x, y := PixelsToMeters(px, py, zoom)
	return MetersToLatLon(x, y)
}

// Returns a tile covering region in given pixel coordinates
func PixelsToTile(px, py float64) (int, int) {
	tileX := int(math.Floor(px / tileSize))
	tileY := int(math.Floor(py / tileSize))
	return tileX, tileY
}

// Returns tile for given mercator coordinates
func MetersToTile(x, y float64, zoom int) (int, int) {
	px, py := MetersToPixels(x, y, zoom)
	return PixelsToTile(px, py)
}